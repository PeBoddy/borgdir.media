var check = function () {
    var user = document.getElementById('user').value;
    var mail = document.getElementById('mail').value;
    var psw = document.getElementById('psw').value;
    var pswr = document.getElementById('pswr').value;

    if(user != "" && mail != "" && psw != "" && pswr != "") {
        if (psw == pswr) {
            document.getElementById('message').style.color = 'green';
            document.getElementById('message').innerHTML = 'matching';
        } else {
            document.getElementById('message').style.color = 'red';
            document.getElementById('message').innerHTML = 'not matching';
            document.getElementById('btnRegister').hidden;
        }
    }
    if(user != "" && mail != "" && document.getElementById('message').innerHTML == 'matching') {
        document.getElementById('btnRegister').disabled = false;
    }
    if(user  == "" || mail == "" || document.getElementById('message').innerHTML == 'not matching') {
        document.getElementById('btnRegister').disabled = true;
    }

};