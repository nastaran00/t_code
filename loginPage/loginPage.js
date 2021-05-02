 
 function endLogIn(){
        //var em=document.getElementById("emailInputId").value
        //reg=/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        //var pass=document.getElementById("passwordInputId").value
        //if(reg.test(em) && pass.length>=6){
             //inaj bayad etelaat da data base ham check beshe!!
             //alert("here")
            //window.location.assign("../homepage/homepage.html")
            //alert("next")
        //}
        //else if(!reg.test(em)){
          //  document.getElementById("incorrectmsg").innerHTML="wrong email!"
            //alert("email is not correct!")
        //}
        //else if(pass.length<6){
          //  document.getElementById("incorrectmsg").innerHTML="password length should be more than 5 !"
            //alert("password should be more than 5 characters!")
        //}




        
    const form=document.getElementById("loginForm")
    formData=new FormData(form)
    req={
        method:'POST',
       // headers: {'Content-Type': 'multipart/form-data'},
        body: formData
    }
    const url="http://localhost:8001/loginPage"
    try{
        fetch(url,req)
        .then(function(response){
            if(response.ok){
                //alert((req['body']).get("texttw"))
                //alert(req['body'].get("imagetw"))
                window.location.assign("../homepage/homepage.html")
                
            }
        })
    }catch(error){

    }
}
function signUpfromlogin(){
    window.location.assign("../signupPage/signupPage.html")
}
