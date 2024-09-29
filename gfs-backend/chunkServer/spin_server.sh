#The script would spin up n chunkservers, they are opened using tmux.
n=3
go_server_cmd="go run server.go"

tmux new-session -d -s myservers


for ((i=1; i<=n; i++)); do
  tmux new-window -t myservers:$i -n "server-$i" "$go_server_cmd"
  echo "Started Go server $i in tmux window server-$i"
done


tmux attach -t myservers