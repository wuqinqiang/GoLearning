package models

import "time"

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id !=0{
		valid=true
	}
	return
}

func (session *Session) DeleteByUUID()(err error)  {
	statement:="delete from sessions where uuid=?"
	stmt,err:=Db.Prepare(statement)
	if err !=nil{
		return
	}
	defer stmt.Close()
	_,err=stmt.Exec(session.Uuid)
	return
}

func (session *Session) User()(user User,err error)  {
	user=User{}
	err=Db.QueryRow("select id,uuid,name,email,created_at from users where id=?",session.UserId).
		Scan(&user.Id,&user.Uuid,&user.Email,&user.Name,&user.CreatedAt)
	return
}
func SessionDeleteAll()(err error)  {
	statement:="delete from sessions"
	_,err=Db.Exec(statement)
	return
}
