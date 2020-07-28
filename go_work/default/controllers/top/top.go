package top

import (
    _ "fmt"
    "os"
    "net/http"
    "github.com/gin-gonic/gin"

    "default/models/db"
    "default/models/db/objects/contentsInfo"
)

func Index( c *gin.Context ) {

    // DB Access Sample
    db := dbManager.ConnectDatabase( os.Getenv("DB_NAME") )
    contentsInfoDatas := contentsInfo.FindAll( db )

    //-----------------------------
    // find  use case
    //-----------------------------
    // findDatas := contentsInfo.Find( db , "where id=4" )

    //-------------------
    // update use case
    //-------------------
    // changedValue := contentsInfo.Params {
    //     ContentKey : "EFG" ,
    // }
    // contentsInfo.Update( db , "where content_key like 'XX%'" , changedValue )

    //-------------------
    // delete use case
    //-------------------
    // contentsInfo.Delete( db , "where id=8" )

    //-------------------
    // insert use case
    //-------------------
    // cip := contentsInfo.Params {
    //     ContentKey : "ABC" ,
    // }
    // contentsInfo.Insert( db , cip )

    c.HTML(http.StatusOK, "top.html", gin.H{
        "Values": contentsInfoDatas ,
    })
}
func Admin( c *gin.Context ) {
    c.HTML(http.StatusOK, "admin.html", gin.H{
        "value": "page-top-admin",
    })
}