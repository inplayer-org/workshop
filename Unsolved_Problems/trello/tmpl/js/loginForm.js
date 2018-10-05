window.onload = AddEventListeners;

function AddEventListeners(){
  login = document.getElementById('logInHeader')
  signUp = document.getElementById("signUpHeader");
  login.addEventListener("click",testFunction.bind(this,"logInHeader"));
  signUp.addEventListener("click",testFunction.bind(this,"signUpHeader"));
}

function testFunction(ob){
  if(ob=="logInHeader"){
    document.getElementById('logInHeader').classList.add("selectedHeader");
    document.getElementById('logInForm').classList.remove("hideForm");
    document.getElementById('signUpHeader').classList.remove("selectedHeader");
    document.getElementById('signUpForm').classList.add("hideForm");
  }else{
    document.getElementById('logInHeader').classList.remove("selectedHeader");
    document.getElementById('logInForm').classList.add("hideForm");
    document.getElementById('signUpHeader').classList.add("selectedHeader");
    document.getElementById('signUpForm').classList.remove("hideForm");
  }
}
