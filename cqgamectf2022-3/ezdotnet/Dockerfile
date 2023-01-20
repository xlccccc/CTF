FROM mcr.microsoft.com/dotnet/aspnet:6.0
COPY app /app
COPY start.sh /start.sh
RUN chmod +x /start.sh
WORKDIR /app
RUN chmod u+s /bin/cat
RUN useradd dotnet
RUN chown dotnet:dotnet -R /app
ENTRYPOINT ["/start.sh"]
