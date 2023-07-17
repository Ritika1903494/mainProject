package controller

import (
	"fmt"
	"encoding/json"
	db "mainproject/Database"
	model "mainproject/model"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)


var Db = db.InitDb()

func CreateNote(c *gin.Context) {
	fmt.Println("12345")
	 note:= model.Notebook{}
	jsondata,_:= c.GetRawData()
	fmt.Println("jsondata",string(jsondata))
	bytes:=[]byte(jsondata)
	fmt.Println("bytes",bytes)
	err := json.Unmarshal(bytes,&note)
         if err != nil {
              fmt.Println("Can not unmarshal the byte array")
            return
		}  
	fmt.Println("note",note.Name)
    result := Db.Create(&note); 
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": result.Error})
        return
    }
    c.JSON(http.StatusCreated, &note)
	return
}

func GetNotes(c *gin.Context) {
	fmt.Println("get..")
	var notes []model.Notebook
	 result := Db.Find(&notes); 
	 if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": result.Error})
        return
    }
	  c.JSON(http.StatusOK,&notes)
	  return
}


func CreateSection(c *gin.Context) {
	fmt.Println("section")
	 section:= model.Section{}
	jsondata,_:= c.GetRawData()
	fmt.Println("jsondata",string(jsondata))
	bytes:=[]byte(jsondata)
	fmt.Println("bytes",bytes)
	err := json.Unmarshal(bytes,&section)
         if err != nil {
              fmt.Println("Can not unmarshal the byte array")
            return
		}  
	fmt.Println("section",section.Name,section.Note_ID)
    result := Db.Create(&section); 
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": result.Error})
        return
    }
    c.JSON(http.StatusCreated, &section)
	return
}
func GetSections(c *gin.Context) {
	note_id,_:= strconv.Atoi(c.Param("note_id"))
	fmt.Println("section..")
	var sections []model.Section
	 result := Db.Where("note_id = ?",note_id).Find(&sections) 
	 if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": result.Error})
        return
    }
	  c.JSON(http.StatusOK,&sections)
	  return
}

func CreatePage(c *gin.Context) {
	fmt.Println("page")
	 page:= model.Page{}
	jsondata,_:= c.GetRawData()
	fmt.Println("jsondata",string(jsondata))
	bytes:=[]byte(jsondata)
	fmt.Println("bytes",bytes)
	err := json.Unmarshal(bytes,&page)
         if err != nil {
              fmt.Println("Can not unmarshal the byte array")
            return
		}  
	fmt.Println("page",page.Name,page.Sec_ID)
    result := Db.Create(&page); 
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": result.Error})
        return
    }
    c.JSON(http.StatusCreated, &page)
	return
}

func GetPages(c *gin.Context) {
	sec_id,_:= strconv.Atoi(c.Param("section_id"))
	fmt.Println("page..")
	var pages []model.Page
	 result := Db.Where("sec_id = ?",sec_id).Find(&pages) 
	 if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": result.Error})
        return
    }
	  c.JSON(http.StatusOK,&pages)
	  return
}

func CreateContent(c *gin.Context) {
	fmt.Println("Content")
	 content:= model.Content{}
	jsondata,_:= c.GetRawData()
	fmt.Println("jsondata",string(jsondata))
	bytes:=[]byte(jsondata)
	fmt.Println("bytes",bytes)
	err := json.Unmarshal(bytes,&content)
         if err != nil {
              fmt.Println("Can not unmarshal the byte array")
            return
		}  
	fmt.Println("content",content.Page_Content,content.Pg_ID)
    result := Db.Create(&content); 
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": result.Error})
        return
    }
    c.JSON(http.StatusCreated, &content)
	return
}

func GetContent(c *gin.Context) {
	pg_id,_:= strconv.Atoi(c.Param("page_id"))
	fmt.Println("content..")
	var content [] model.Content
	 result := Db.Where("pg_id = ?",pg_id).Find(&content) 
	 if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error":"record dont found"})
        return
    }
	  c.JSON(http.StatusOK,&content)
	  return
}

func UpdateContent(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("content_id"))
	body:= model.Content{}
	jsondata,_:= c.GetRawData()
	fmt.Println("jsondata",string(jsondata))
	bytes:=[]byte(jsondata)
	fmt.Println("bytes",bytes)
	err := json.Unmarshal(bytes,&body)
         if err != nil {
              fmt.Println("Can not unmarshal the byte array")
            return
		}  
	fmt.Println("body",body.Page_Content)
     
	var content model.Content

    if result := Db.First(&content,id); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }
    content.Page_Content=body.Page_Content
    Db.Save(&content)
    c.JSON(http.StatusOK,content)
}