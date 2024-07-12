echo "Please enter image version"
read iamgeVersion

docker build -t demo-http-server .

# 镜像tap
docker tag demo-http-server qvbilam/http-server:"$iamgeVersion"-alpine3.15

# 提交镜像版本
docker commit -a "qvbilam" -m "go http server" qvbilam/http-server:"$iamgeVersion"-alpine3.15 -p

# 登陆
docker login

# 推送
docker push qvbilam/http-server:"$iamgeVersion"-alpine3.15