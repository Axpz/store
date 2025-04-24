package main

func localStorage() {
	// var cfg config.Config

	// cfg.Storage.Path = "/tmp"
	// cfg.Storage.Type = "local"

	// // 初始化存储引擎
	// store, err := storage.New(&cfg)
	// if err != nil {
	// 	log.Fatalf("初始化存储引擎失败: %v", err)
	// }

	// // 存储一些数据
	// examples := map[string]any{
	// 	"user:1": map[string]any{
	// 		"name":  "张三",
	// 		"age":   25,
	// 		"email": "zhangsan@example.com",
	// 	},
	// 	"user:2": map[string]any{
	// 		"name":  "李四",
	// 		"age":   30,
	// 		"email": "lisi@example.com",
	// 	},
	// }

	// fmt.Println("存储数据...")
	// for key, value := range examples {
	// 	if err := store.Put(key, value); err != nil {
	// 		log.Fatalf("存储数据失败: %v", err)
	// 	}
	// 	fmt.Printf("已存储: %s\n", key)
	// }

	// // 读取数据
	// fmt.Println("\n读取数据...")
	// for key := range examples {
	// 	value, err := store.Get(key)
	// 	if err != nil {
	// 		log.Fatalf("读取数据失败: %v", err)
	// 	}
	// 	fmt.Printf("%s: %v\n", key, value)
	// }

	// // 删除数据
	// fmt.Println("\n删除数据...")
	// if err := store.Delete("user:1"); err != nil {
	// 	log.Fatalf("删除数据失败: %v", err)
	// }
	// fmt.Println("已删除: user:1")

	// // 验证删除
	// _, err = store.Get("user:1")
	// if err != nil {
	// 	fmt.Printf("验证删除成功: %v\n", err)
	// }
}

func main() {
	localStorage()
}
