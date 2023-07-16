
function sendData(event){
    event.preventDefault();
    const name = document.getElementById('formControlName');
    const email = document.getElementById('formControlEmail');
    const phone = document.getElementById('formControlphone');
    const subject = document.getElementById('formControlSubject');
    const msg = document.getElementById('exampleFormControlTextarea1');

    if(name.value === ''){
        name.setAttribute('required', true);
    }else if(email.value === ''){
        email.setAttribute('required', true);
    }else if(phone.value === ''){
        phone.setAttribute('required', true);
    }else if(subject.value === ''){
        subject.setAttribute('required', true);
    }else if(msg.value === ''){
        msg.setAttribute('required', true);
    }else{
        const mailTO = 'devfajarmalik@gmail.com';

        let a = document.createElement('a')
        a.href = `mailto:${mailTO}?subject=${subject.value}&body=Halo nama saya :${name.value},message:${msg.value},silahkan kontak saya di nomor berikut: ${phone.value}/email:${email.value}`
        a.click()
    }
}

// download resume
const link = document.querySelector('.download');
link.addEventListener('click', function(){
    link.download = 'image/HTML.png';
});



