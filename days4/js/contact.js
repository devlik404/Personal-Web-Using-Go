
function sendData(event){
    event.preventDefault();
    const name = document.getElementById('name');
    const email = document.getElementById('email');
    const phone = document.getElementById('tlp');
    const subject = document.getElementById('opt');
    const msg = document.getElementById('msg');

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
