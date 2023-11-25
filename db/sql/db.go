package db


// 增删改查

// Valuer 和 Scanner 接口

// 事务和隔离级别


// Prepare Statement

// sqlmock 入门

// SQL 编程面试要点


// 

/*

连接

	类型
		直接连接 / Conn
		预编译 / Stmt
		事务 / Tx
			
操作
	处理返回数据的几种方式
		Exec -> Result
		Query -> Rows (多条记录)
		QueryRow -> Row（简化单条记录，字段）
 
*/


// type Base struct {
// 	ID 	       int        `json:"id"`
// 	CreatedAt  time.Time  `json:"created_at"`
// 	UpdatedAt  time.Time  `json:"updated_at"`
// }

// type User struct {
// 	Base
// 	Email      string     `json:"email"`
// 	Name       string     `json:"first_name,omitempty"`
// 	Password   string     `json:"-"`
// 	Active     int        `json:"active"`
// }
	
// func (u *User) TableName() string {
// 	return "user"
// }






// 	// Insert a new user
// 	{
// 	// 	const q = `
// 	// 	`
// 		// 	const q = 
// // 	`
// // 	SELECT
// // 		id, email, name, password, user_active, created_at, updated_at
// // 	FROM
// // 		users
// // 	WHERE 
// // 		email = ?
// // 	`

// 		username := "admin"
// 		password := "123456"
// 		createdAt := time.Now()

// 		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		id, err := result.LastInsertId()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(id)
// 	}

// 	// Query a single user
// 	{ 
		
// 		const query = 
// 		`
// 		SELECT 
// 			id, username, password, created_at 
// 		FROM 
// 			user 
// 		WHERE 
// 			id = ?
// 		`

// 		var (
// 			id        int
// 			username  string
// 			password  string
// 			createdAt time.Time
// 		)

// 		if err := db.QueryRow(query, 1).Scan(
// 			&id, 
// 			&username, 
// 			&password, 
// 			&createdAt); err != nil {
			
// 			log.Fatal(err)
// 		}

// 		fmt.Println(id, username, password, createdAt)
// 	}

// 	// Query all users
// 	{ 
// 		type user struct {
// 			id        int
// 			username  string
// 			password  string
// 			createdAt time.Time
// 		}

// 		rows, err := db.Query(`SELECT id, username, password, created_at FROM user`)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer rows.Close()

// 		var users []user
// 		for rows.Next() {
// 			var u user

// 			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			users = append(users, u)
// 		}
// 		if err := rows.Err(); err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Printf("%#v", users)
// 	}

// 	// Delete a single user
// 	{
// 		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }