function createUser(){
  //alert("dd")
  //window.location.assign("../homepage/homepage.html")

 // httpRequest= new XMLHttpRequest();
// httpRequest.load= function() {
 //          alertContents(this.responseText)
//};
httpRequest.open("GET", "http://localhost:8001/homepage/homepage.html", true);
//httpRequest.send();
}
async function get_objects(){ ///get all available objects

  document.getElementById("createAccount").submit();


    var base_url = "http://localhost:8001/homepage/";
    ///////////send request for a _specified number_->var number =20 (in this case)
    url = base_url //+ num + "/";////create new url
    var res_fetch = fetch(url)
    .then( function(response){
        if (response.ok) {
          alert(response.json())
            return response.json()
          } else if(response.status === 404) {////////////Error hanling for API ////
            return Promise.reject('error 404')
          } else {
            return Promise.reject('some other error: ' + response.status)
          }
    });
    await res_fetch.then(function(json){

    
        if (!json){
            return;
        }
        console.log("json :))) ",json);

    }).catch(error => {  ////////////Error hanling for API////
        console.log('error is :', error,"(url:",url,")")
        //window.alert('error is :', error,"(url:",url,")")
    });
        
//await init_();

}
function test1(){
  alert("tesss")
}
//get_objects();