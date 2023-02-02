<?php
if (isset($_POST['jpgraw'])) {
    $jpgraw = $_POST['jpgraw'];
    $raw = file_get_contents($jpgraw);
    $img = imagecreatefromstring($raw);
    imagejpeg($img,'easy.jpg');
    if($img != false) {
        file_put_contents(md5($_POST['filename']).'.txt',$raw);
    }
}
//🎉新年快乐
?>