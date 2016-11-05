deps:
	wget https://github.com/imperavi/kube/archive/v6.0.1.zip -O kube.zip
	unzip -o kube.zip
	mv kube-6.0.1/dist/css/* assets/css
	mv kube-6.0.1/dist/js/* assets/js
