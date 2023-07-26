const promises = new Promise((Rlv,Rjc ) => {
    const xhr = new XMLHttpRequest()

    xhr.open('GET','https://api.npoint.io/08e5931c464bf7a806db',true)

    xhr.onload = () =>{
        // http Request success status response code 
    if(xhr.status === 200){
        Rlv(JSON.parse(xhr.responseText))
    }else if(xhr.status === 400){
        Rjc("error Request :(")
    }
    }
    xhr.onerror = () => {
    Rjc("NetWork Error")
    }
    xhr.send()
    
})


// promises.then(value => console.log(value)).catch(reason => console.log(reason))

let RatingTestimonial = []

async function getData (){
        try{
    const result = await promises
    RatingTestimonial = result
    allRating()
        }catch(error){
            console.log(error)
        }
    }
    // memanggil funsi get data 
getData()
// menampilkan semua rating
    function allRating(){
        let ratingTestimoni = ""
        RatingTestimonial.forEach(cardRate => ratingTestimoni += `
        <div class="card">
            <img src="${cardRate.image}" alt="Avatar" style="width:100%">
            <div class="container">
            <i>"${cardRate.quote}"</i> 
            <div class="starAutor">
            <h4><b>${cardRate.autor}</b></h4> 
            <span><i class="fa fa-star" aria-hidden="true">${cardRate.rating}</i></span>
            </div>
            </div>
            </div>`
        )
        document.getElementById('container-card').innerHTML = ratingTestimoni;
    }

    //menyeleksi setiap rating
    function StarRating(Star){
    let filterRate = ""
    
    RatingTestimonial.filter(star => star.rating === Star)
        .forEach(cardRate => filterRate += `
        <div class="card">
            <img src="${cardRate.image}" alt="Avatar" style="width:100%">
            <div class="container">
            <i>"${cardRate.quote}"</i> 
            <div class="starAutor">
            <h4><b>${cardRate.autor}</b></h4> 
            <span><i class="fa fa-star" aria-hidden="true">${cardRate.rating}</i></span>
            </div>
            </div>
            </div>`
            )
        document.getElementById('container-card').innerHTML = filterRate;
    }

