package main
import(
    "fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
)
//数据库配置
cinst(
	userName = "root"
	password = "smmdb2016"
	ip = "182.254.211.181"
	port = "3306"
	dbName = "lme_price_02"
)
//数据库连接池
var DB *sql.DB

func InitDB(){
	//构建连接
    path :=strings.join([]string{userName,":",password,"@tcp(",ip,":",port,")/",dbName,"?charset=utf8"},"")
	//打开数据库，导入驱动
	DB,_=sql.Open("mysql",path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err :=DB.Ping();err !=nil{
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connect success")

}

func main(){
    InitDB()
	//调用接口
	resp, err :=http.Get("https://www.lme.com/api/trading-data/day-delayed?datasourceId=2a431297-6620-4ba7-a991-8335423f994b")
	//检查接口是否通了
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	//输出结果
	body,_ :=ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	var res result
	_ = json.Unmarshal(body,&res)
	fmt.Printf("%#v",res)

}
