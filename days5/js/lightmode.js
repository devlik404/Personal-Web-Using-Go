
let bgcolor = document.querySelector('.toggler')
let mode = document.getElementById('lightmode');
mode.addEventListener('click', () => {

    if(mode.checked === true){
        document.body.classList.add('light');
        if(mode.checked === true){
            bgcolor.classList.remove('bgnColor');
            
        }
    }else{
        document.body.classList.remove('light');
        bgcolor.classList.add('bgnColor');
    }

});