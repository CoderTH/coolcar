package main

func main() {
	//c, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	//if err != nil {
	//	​     panic(err)
	//}
	//​
	//ctx := context.Background()	//c, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	//if err != nil {
	//	​     panic(err)
	//}
	//​
	//ctx := context.Background()
	//resp, err := c.ContainerCreate(ctx, &container.Config{
	//	Image: "mongo:4.4",
	//	ExposedPorts: nat.PortSet{
	//		"27017/tcp": {},
	//	},
	//}, &container.HostConfig{
	//	PortBindings: nat.PortMap{
	//		"27017/tcp": []nat.PortBinding{
	//			{
	//				HostIP:   "127.0.0.1",
	//				HostPort: "27018",
	//			},
	//		},
	//	},
	//}, nil, "")
	//err = c.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	//if err != nil {
	//	panic(err)
	//}
}
