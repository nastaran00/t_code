

function loginFunc(){
    var base_url = "http://localhost:8001/loginPage";
    ///////////send request for a _specified number_->var number =20 (in this case)
        url = base_url //+ num + "/";////create new url
        var res_fetch = fetch(url)
        .then( function(response){
            if (response.ok) {
                return response.json()
              } else if(response.status === 404) {////////////Error hanling for API ////
                return Promise.reject('error 404')
              } else {
                return Promise.reject('some other error: ' + response.status)
              }
        });
        
            window.location.assign("../loginPage/loginPage.html")
    
        
}
function signupFunc(){
    //window.location.assign("../signupPage/signupPage.html")
    //httpRequest.open("GET", "http://localhost:8002/signupPage", true);
    get_objects();

}

async function get_objects(){ ///get all available objects
    var base_url = "http://localhost:8001/signupPage";
    ///////////send request for a _specified number_->var number =20 (in this case)
        url = base_url //+ num + "/";////create new url
        var res_fetch = fetch(url)
        .then( function(response){
            if (response.ok) {
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
            window.location.assign("../signupPage/signupPage.html")
    
        }).catch(error => {  ////////////Error hanling for API////
            console.log('error is :', error,"(url:",url,")")
            //window.alert('error is :', error,"(url:",url,")")
        });
            
    //await init_();
    
    }