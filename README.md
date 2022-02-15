# updownloader

一个支持命令行、PC端、移动端上传、下载文字或者文件的前后端分离网站。

CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' .


location ^~/updown {
		alias /home/ubuntu/updownloader/dist;
		index index.html index.htm;
		try_files $uri $uri/ /updown/index.html;
	}