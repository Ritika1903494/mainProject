package model


type Notebook struct {
	Notebook_ID   int   `json:"note_id" gorm:"primaryKey;autoIncrement;not null"`
	Name  string  `json:"note_name" binding:"required"`
	UserId int `json:"user_id"`
	Sections []Section `json:"note_fk" gorm:"foreignKey:Note_ID;references:Notebook_ID"`
}

type Section struct {
	Section_ID    int `json:"section_id" gorm:"primaryKey;autoIncrement;not null"`
	Name  string  `json:"section_name" binding:"required"`
	Note_ID int  `json:"note_id"` 
	Pages []Page `json:"sec_fk" gorm:"foreignKey:Sec_ID;references:Section_ID"`
}

type Page struct{
	Page_ID    int `json:"page_id" gorm:"primaryKey;autoIncrement;not null"`
	Name  string  `json:"page_name" binding:"required"`
	Sec_ID int  `json:"section_id"`
	Content Content `json:"page_fk" gorm:"foreignKey:Pg_ID;references:Page_ID"`
}

type Content struct {
	Content_Id  int `json:"content_id" gorm:"primaryKey;autoIncrement;not null"`
	Page_Content string `json:"page_content"`
	Pg_ID int  `json:"page_id"`
}