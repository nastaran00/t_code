//import renderList from "../App.js"
//import React,{useEffect,useState} from 'react'
var userId
//const home=require("./App")
function logout() {
    
    //var a=document.getElementById("editpage")
    document.getElementById("logoutButton").style.color="blue"
    document.getElementById("homeIcon").style.color="black"
    document.getElementById("editprofileIcon").style.color="black"

        



       //irad dare
    req={
        method:'GET',
       // headers: {'Content-Type': 'multipart/form-data'},
        
    }
    const url="http://localhost:8001/logout/"+userId //eslah konam
    try{
        fetch(url,req)
        .then(function(response){
            if(response.ok){
                alert("logout")
                window.location.assign("../FirstPage/twitter.html")

                
            }
        })
    }catch(error){

    }


    
}
function changeProfile() {
    
    const text=document.getElementById("newNamee").value
    alert(text)
    //const form=document.getElementById("editForm")
    //formDatachange=new FormData(form)
    req={
        method:'PUT',
       // headers: {'Content-Type': 'multipart/form-data'},
        body: text
    }
    const url="http://localhost:8001/homepage/edit/"+userId 
    try{
        fetch(url,req)
        .then(function(response){
            if(response.ok){
                alert(response.newName)
                

                
            }
        })
    }catch(error){

    }


}
function setid(i){
    userId=i
    alert("yess"+userId)
}
function goToHome(){
    document.getElementById("editprofileIcon").style.color="black"
    document.getElementById("homeIcon").style.color="blue"
    document.getElementById("logoutButton").style.color="black"
    var b=document.getElementById("centerhome")
   
        if(b.style.display=="none"){
            //alert("here")
            b.style.display="block"
            document.getElementById("editpage").style.display="none"
        }
}
function goToEdit(){

    var a=document.getElementById("editpage")
    document.getElementById("editprofileIcon").style.color="blue"
    document.getElementById("homeIcon").style.color="black"
    document.getElementById("logoutButton").style.color="black"
    

        if(a.style.display=="none"){
            //alert("here")
            a.style.display="block"
            document.getElementById("centerhome").style.display="none"
        }
}
function tweetIt(){//event ro hazf konam!!
    
    const text=document.getElementById("texttweet").value
    alert(text)
    const form=document.getElementById("f")
    formData=new FormData(form)
    j={
        tex:text
    }
    req={
        method:'PUT',
        //headers: {'Content-Type': 'application/json'},
        body: formData
    }
    const url="http://localhost:8001/homepage/tweet/"+userId //id ham mikhad
    try{
        fetch(url,req)
        .then(function(response){
            if(response.ok){
               getAlltweetsForShow()

                innerdiv=document.createElement("div")
                innerdiv.style.display="grid"
                innerdiv.style.gridTemplateColumns=" 40% 60% "
                tweetText=document.createTextNode(text)
                innerdiv.appendChild(tweetText)
                tweetImage=document.createElement("img")
               // tweetImage.src=req['body'].get("imagetw")//aks!!!
                innerdiv.appendChild(tweetImage)
                


                innerdiv.style.borderBottom="1px solid blue"
                innerdiv2=document.createElement("div")
                innerdiv2.id="innerdiv2ListlikeDeleteRetweet"
                innerdiv2.style.display="grid"
                innerdiv2.style.gridTemplateColumns=" 30% 30% 40%"
                innerdiv2.style.borderBottom="1px solid black"
                innerdiv2.id="inner2"
                grid1=document.createElement("div")
                //grid1.id="grid1Like"+res.tweet_id[i]
               // grid1.id=res.tweet_id[i]
                grid1.onclick=function(){
                    likeTweet()
                }
                g1text=document.createTextNode("Like")
                grid1.style.fontSize="xx-small"
                grid1.appendChild(g1text) 

                grid2=document.createElement("div")
                //grid2.id="grid2Retweet"+res.tweet_id[i]
               // grid2.id=res.tweet_id[i]
                grid2.onclick=function(){
                    ReTweet()
                }
                g2text=document.createTextNode("Retweet")
                grid2.style.fontSize="xx-small"
                grid2.appendChild(g2text)

                grid3=document.createElement("div")
                //grid3.id="grid3Delete"+res.tweet_id[i]
               // grid3.id=res.tweet_id[i]

                grid3.onclick=function(){
                    DeleteTweet(this)
                }
                g3text=document.createTextNode("Delete")
                grid3.style.fontSize="xx-small"
                grid3.appendChild(g3text)

                innerdiv2.appendChild(grid1)
                innerdiv2.appendChild(grid2)
                innerdiv2.appendChild(grid3)

                listofT=document.getElementById("listOfTweetSection")
                listofT.appendChild(innerdiv)
                listofT.appendChild(innerdiv2)

                //var script=document.createElement("script")
                //script.src="../App.js"
                //document.head.appendChild(script)
                //$.getscript("../App.js",App(formData))
                //System.import('../signupPage/signupPage.js').then(()=>{
                    //test1()
                   // t()
                //})
                //$.getscript("../signupPage/signupPage.js",test1())
                alert("no")
            }
        })
    }catch(error){

    }
}

function getAlltweetsForShow(){//??eslah mikhad
alert("again")
    req={
        method:'GET',
       // headers: {'Content-Type': 'multipart/form-data'},
      
    }
    const url="http://localhost:8001/homepage/"+userId //id ham mikhad
    try{
        fetch(url,req).then(response=>response.json())
        .then(function(response){
            alert("innnn")
            res=JSON.stringify(response)
            alert("here"+res)
            if(res.msg==="success"){
                tweets=res.tweets
                alert(tweets)
                for(var i=0;i<tweets.length;i++){ //inja tweet hasho az db migire neshoon mide

                    innerdiv=document.createElement("div")
                    innerdiv.style.display="grid"
                    innerdiv.style.gridTemplateColumns=" 40% 60% "// baraye matn va aks
                    //innerdiv.id="innerdiv"+res.tweet_id[i]
                    innerdiv.id=res.tweet_id[i]
                    tweetText=document.createTextNode(text)
                    innerdiv.appendChild(tweetText)
                    tweetImage=document.createElement("img")
                   // tweetImage.src=req['body'].get("imagetw")//aks!!!
                    innerdiv.appendChild(tweetImage)
                    
    
    
                    innerdiv.style.borderBottom="1px solid blue"
                    innerdiv2=document.createElement("div")
                    innerdiv2.id="innerdiv2ListlikeDeleteRetweet"
                    innerdiv2.style.display="grid"
                    innerdiv2.style.gridTemplateColumns=" 30% 30% 40%"
                    innerdiv2.style.borderBottom="1px solid black"
                    innerdiv2.id="inner2"
                    grid1=document.createElement("div")
                    //grid1.id="grid1Like"+res.tweet_id[i]
                    grid1.id=res.tweet_id[i]
                    grid1.onclick=function(){
                        likeTweet()
                    }
                    g1text=document.createTextNode("Like")
                    grid1.style.fontSize="xx-small"
                    grid1.appendChild(g1text) 
    
                    grid2=document.createElement("div")
                    //grid2.id="grid2Retweet"+res.tweet_id[i]
                    grid2.id=res.tweet_id[i]

                    grid2.onclick=function(){
                        ReTweet()
                    }
                    g2text=document.createTextNode("Retweet")
                    grid2.style.fontSize="xx-small"
                    grid2.appendChild(g2text)
    
                    grid3=document.createElement("div")
                    //grid3.id="grid3Delete"+res.tweet_id[i]
                    grid3.id=res.tweet_id[i]
                    grid3.onclick=function(){
                        DeleteTweet(this)
                    }
                    g3text=document.createTextNode("Delete")
                    grid3.style.fontSize="xx-small"
                    grid3.appendChild(g3text)
    
                    innerdiv2.appendChild(grid1)
                    innerdiv2.appendChild(grid2)
                    innerdiv2.appendChild(grid3)
    
                    listofT=document.getElementById("listOfTweetSection")
                    listofT.appendChild(innerdiv)
                    listofT.appendChild(innerdiv2)


                }
            }
            

        })
    }catch(error){

    }

}

function DeleteTweet(elem){
alert("delete")


    document.getElementById("inner2").onclick=function(ev){
        alert("hi")
        alert(ev.target.id)
        //const parentOfthisDeleteDiv=elem.parentNode.id
const tweet=ev.target.id //bayad meghdar dehi she
   
data={
    id:userId,
    tweetid:tweet
}
req={
    method:'DElETE',
   // headers: {'Content-Type': 'multipart/form-data'},
    body: data
}
//  /homepage/tweet/:id/:tweet
const url="http://localhost:8001/homepage/tweet/"+userId+"/"+tweet 
try{
    fetch(url,req).then(response=>response.json())
    .then(function(response){//dobare baraye namayesh migirim
        if(response.ok){
            alert(response.json()+userId)


            res=JSON.stringify(response)
            alert("here"+res)
            if(res.msg==="success"){
                tweets=res.tweets
                alert(tweets)
                for(var i=0;i<tweets.length;i++){ //inja tweet hasho az db migire neshoon mide

                    innerdiv=document.createElement("div")
                    innerdiv.style.display="grid"
                    innerdiv.style.gridTemplateColumns=" 40% 60% "// baraye matn va aks
                    //innerdiv.id="innerdiv"+res.tweet_id[i]
                    innerdiv.id=res.tweet_id[i]
                    tweetText=document.createTextNode(text)
                    innerdiv.appendChild(tweetText)
                    tweetImage=document.createElement("img")
                   // tweetImage.src=req['body'].get("imagetw")//aks!!!
                    innerdiv.appendChild(tweetImage)
                    
    
    
                    innerdiv.style.borderBottom="1px solid blue"
                    innerdiv2=document.createElement("div")
                    innerdiv2.id="innerdiv2ListlikeDeleteRetweet"
                    innerdiv2.style.display="grid"
                    innerdiv2.style.gridTemplateColumns=" 30% 30% 40%"
                    innerdiv2.style.borderBottom="1px solid black"
                    innerdiv2.id="inner2"
                    grid1=document.createElement("div")
                    //grid1.id="grid1Like"+res.tweet_id[i]
                    grid1.id=res.tweet_id[i]
                    grid1.onclick=function(){
                        likeTweet()
                    }
                    g1text=document.createTextNode("Like")
                    grid1.style.fontSize="xx-small"
                    grid1.appendChild(g1text) 
    
                    grid2=document.createElement("div")
                    //grid2.id="grid2Retweet"+res.tweet_id[i]
                    grid2.id=res.tweet_id[i]

                    grid2.onclick=function(){
                        ReTweet()
                    }
                    g2text=document.createTextNode("Retweet")
                    grid2.style.fontSize="xx-small"
                    grid2.appendChild(g2text)
    
                    grid3=document.createElement("div")
                    //grid3.id="grid3Delete"+res.tweet_id[i]
                    grid3.id=res.tweet_id[i]
                    grid3.onclick=function(){
                        DeleteTweet(this)
                    }
                    g3text=document.createTextNode("Delete")
                    grid3.style.fontSize="xx-small"
                    grid3.appendChild(g3text)
    
                    innerdiv2.appendChild(grid1)
                    innerdiv2.appendChild(grid2)
                    innerdiv2.appendChild(grid3)
    
                    listofT=document.getElementById("listOfTweetSection")
                    listofT.appendChild(innerdiv)
                    listofT.appendChild(innerdiv2)


                }
            }
            
            

        }
    })
}catch(error){

}
   



    }
}
function ReTweet(){
alert("retweet")





}
function likeTweet(){ // bayad eslah she, bayad dota id bede o dar vaghe bedoonim che tweeti ro

    alert("like")

    document.getElementById("inner2").onclick=function(ev){
        //alert("hi")
        alert(ev.target.id)
        //const parentOfthisDeleteDiv=elem.parentNode.id
const tweet=ev.target.id //bayad meghdar dehi she
   

//const parentOfthisDeleteDiv=elem.parentNode.id
//const idTobeLiked=1 //bayad meghdar dehi she
   
ourdata={
   // id_liked:userId,
    id_to_be_liked:tweet 
}
req={
    method:'POST',
   // headers: {'Content-Type': 'multipart/form-data'},
    body: ourdata
}
//  /homepage/tweet/:id/:tweet
const url="http://localhost:8001/homepage/tweet/"+userId //bayad eslah she
try{
    fetch(url,req)
    .then(function(response){
        if(response.ok){
            alert("like for tweetId"+tweet)
            

        }
    })
}catch(error){

}
    }
}