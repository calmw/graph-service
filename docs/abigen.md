## 生成绑定文件

- 生成绑定go文件命令

#### 生成userLocation合约代码

- abigen --abi ./abi/NewbieTask.json --pkg intoBinding --type NewbieTask --out ./binding/NewbieTask.go
- abigen --abi ./abi/SyncContact.json --pkg intoBinding --type SyncContact --out ./binding/SyncContact.go

