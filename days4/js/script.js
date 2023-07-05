
let FakeD =[];

const Blogs = (e) => {
    e.preventDefault();
    // ambil value dari ID yang ada di tag HTML
    let NameProject= document.getElementById('PName').value;
    let FirstD = document.getElementById('FDate').value;
    let LastD = document.getElementById('LDate').value;
    let Message = document.getElementById('TxtMsg').value;
    let CheckBox1 = document.getElementById('Cbx1').checked;
    let CheckBox2 = document.getElementById('Cbx2').checked;
    let CheckBox3 = document.getElementById('Cbx3').checked;
    let CheckBox4 = document.getElementById('Cbx4').checked;
    let files = document.getElementById('file').files;
   
// Periksa apakah file telah dipilih
    if (files.length > 0) {
        let file = files[0]; 
    // Cek apakah File yang diunggah adalah tipe gambar yang diizinkan (JPG, PNG, atau JPEG)
        const allowedTypes = ['image/jpeg', 'image/png', 'image/jpg'];
        //jika yang di upload adalah tipe gambar  yang diizinkan maka lolos
        if (allowedTypes.includes(file.type)) {
        files = URL.createObjectURL(files[0])

    // buat object dari value yang telah diambil
     let DbLog = {
        NameProject,
        FirstD,
        LastD,
        Message,
        files,
        CheckBox1, 
        CheckBox2, 
        CheckBox3, 
        CheckBox4 
  
        
    }
// push setiap object kedalam variabel kosong
        FakeD.push(DbLog);
// panggil fungsi cards
        Cards();
        }else{
            // jika gagal lulus pengecekan maka akan menampilkan alert
        alert('unggah tipe gambar yang diizinkan (JPG, PNG, atau JPEG)!!')
        }
    }else{
        // jika ID files tidak di inputkan maka akan menampilkan
        alert('Isi gambar terlebih dahulu !!!')
    }
};

function Cards(){
     document.getElementById('container-card').innerHTML='';
    for(let i = 0; i < FakeD.length; i++){
        document.getElementById('container-card').innerHTML += `
            <div class="cards">
            <a href="profile.html" target="_blank"><img src="${FakeD[i].files}" alt="" class="img-pp"></a>
                    <div class="card-name">
                  <h3>${FakeD[i].NameProject}</h3>
                        <p>durasi : ${FakeD[i].FirstD} | ${FakeD[i].LastD}</p>
                        <p>${FakeD[i].Message}</p>
                        <div class="favicon-project">
                        <a href="">  ${FakeD[i].CheckBox1 ? '<img src="../image/icons8-javascript-48.png" alt="JavaScript">' : ''}</a>
                        <a href=""> ${FakeD[i].CheckBox2 ? '<img src="../image/ava-48.png" alt="Java">' : ''}</a>
                        <a href="">   ${FakeD[i].CheckBox3 ? '<img src="../image/golang-48.png" alt="Golang">' : ''}</a>
                        <a href=""> ${FakeD[i].CheckBox4 ? '<img src="../image/python-48.png" alt="Python">' : ''}</a>
                        </div>
                    </div>
                    <div class="btn-edit">
                        <button type="button">edit</button>
                        <button type="reset" >delete</button>
                    </div>
            </div>`
    }
}


