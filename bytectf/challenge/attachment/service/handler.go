package main

import (
	"context"
	"net/url"
	"path"
	"service/kitex_gen/student/management"
	"service/utils"
)

type StudentManagementImpl struct{}

type StudentInfo struct {
	Num    string
	Name   string
	Gender string
}

type TeacherInfo struct {
	Username string
	Password string
	ImgUrl   string
}

var StudentData = make(map[string]StudentInfo, 50)
var TeacherData = make(map[string]TeacherInfo, 5)

func (s *StudentManagementImpl) QueryStudent(ctx context.Context, req *management.QueryStudentRequest) (resp *management.QueryStudentResponse, err error) {
	stu, exist := StudentData[req.Num]
	if !exist {
		return &management.QueryStudentResponse{
			Exist: false,
		}, nil
	}

	resp = &management.QueryStudentResponse{
		Exist:  true,
		Num:    stu.Num,
		Name:   stu.Name,
		Gender: stu.Gender,
	}

	return resp, nil
}

func (s *StudentManagementImpl) InsertStudent(ctx context.Context, req *management.InsertStudentRequest) (resp *management.InsertStudentResponse, err error) {
	_, exist := StudentData[req.Num]
	if exist {
		return &management.InsertStudentResponse{
			Ok:  false,
			Msg: "the num has exists",
		}, nil
	}

	StudentData[req.Num] = StudentInfo{
		Num:    req.Num,
		Name:   req.Name,
		Gender: req.Gender,
	}

	return &management.InsertStudentResponse{
		Ok: true,
	}, nil
}

func (s *StudentManagementImpl) TeacherLogin(ctx context.Context, req *management.TeacherLoginRequest) (resp *management.TeacherLoginResponse, err error) {
	_, exist := TeacherData[req.Username]
	if exist {
		if TeacherData[req.Username].Password == req.Password {
			return &management.TeacherLoginResponse{
				Ok:  true,
				Msg: TeacherData[req.Username].ImgUrl,
			}, nil
		} else {
			return &management.TeacherLoginResponse{
				Ok:  false,
				Msg: "password error",
			}, nil
		}
	}

	return &management.TeacherLoginResponse{
		Ok: false,
	}, nil
}

func (s *StudentManagementImpl) TeacherRegister(ctx context.Context, req *management.TeacherRegisterRequest) (resp *management.TeacherRegisterResponse, err error) {
	_, exist := TeacherData[req.Username]
	if exist && req.Username != "" {
		return &management.TeacherRegisterResponse{
			Ok:  false,
			Msg: "the username has exists",
		}, nil
	}

	var imgUrl string
	if e := utils.IsImgExist(req.ImgUrl); !e {
		fileName, err := utils.DownloadFile(req.ImgUrl)
		if err != nil {
			return &management.TeacherRegisterResponse{
				Ok:  false,
				Msg: "download img error",
			}, nil
		}
		imgUrl = "https://image.local/" + fileName
	} else {
		parsedUrl, err := url.Parse(req.ImgUrl)
		if err != nil {
			return &management.TeacherRegisterResponse{
				Ok:  false,
				Msg: "parse img url error",
			}, nil
		}
		imgUrl = "https://image.local/" + path.Base(parsedUrl.Path)
	}

	TeacherData[req.Username] = TeacherInfo{
		Username: req.Username,
		Password: req.Password,
		ImgUrl:   imgUrl,
	}

	return &management.TeacherRegisterResponse{
		Ok:  true,
		Msg: imgUrl,
	}, nil
}
