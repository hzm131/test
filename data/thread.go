package data

import "time"

type Thread struct{
	Id int
	Uuid string
	Topic string   //主题
	UserId int
	CreatedAt time.Time  //创建时间
}

func Threads() (threads []Thread,err error){
	//向数据库发送一个sql查询，这个查询将返回一个或多个结果
	rows,err := Db.Query("SELECT  id,uuid,topic,user_id,created_at FROM threads ORDER BY created_at DESC")
	if err != nil{
		return
	}
	//遍历行，为每个行分别创建一个Thread结构
	for rows.Next() { //循环执行直到查询返回的所有行都被遍历完毕为止
		th := Thread{}
		//首先使用这个结构去存储行中记录的帖子数据
		if err = rows.Scan(&th.Id,&th.Uuid,&th.Topic,&th.UserId,&th.CreatedAt); err!=nil{
			return
		}
		//然后将存储了帖子数据Thread结构追加到传入的threads切片里面
		threads = append(threads,th)
	}
	rows.Close()
	return
}

//NumReplies方法首先打开一个指向数据库的连接
func (thread *Thread) NumReplies() (count int){
	//通过一条sql语句来取得帖子的数量，并使用传入方法里的count参数来记录这个值，最后返回帖子的数量作为方法的执行结果
	//而模板引擎则使用这个值去代替模板文件中出现的{{.NumReplies}}动作
	rows,err := Db.Query("SELECT  count(*) FROM posts where thread_id=$1",thread.Id)
	if err != nil{
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err!=nil{
			return
		}
	}
	rows.Close()
	return
}