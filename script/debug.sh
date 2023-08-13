# Debug 使用
cd "$(dirname "$0")"
# 1.rpc
# user
go run ../service/user/user.go -f ../service/user/etc/user.yaml &
# video
go run ../service/video/video.go -f ../service/video/etc/video.yaml &
# interaction
go run ../service/interaction/interaction.go -f ../service/interaction/etc/interaction.yaml &
# 2.api
go run ../api/api.go -f ../api/etc/api.yaml &
