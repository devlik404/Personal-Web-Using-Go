//Testimonial Rating menggunakan Higher Order Function

const RatingTestimonial = [
  {
    autor : "- Sanji",
    quote : "Prinsip keadilanku adalah memberi makanan pada yang kelaparan.",
    image : "https://i.pinimg.com/736x/ea/f2/87/eaf287d75e647a43c44be73825fa08ae.jpg",
    rating : 1
  },
  {
    autor : "- Monkey D Luffy",
    quote : "Aku tidak peduli, walaupun aku harus mati untuk mengejar impianku(ngoding)",
    image : "https://th.bing.com/th/id/R.4f0460eb5056b082dbee6ebfc39cd03d?rik=0E8YNEMtN%2bQ4pw&riu=http%3a%2f%2fpm1.narvii.com%2f6362%2f9edfdb1ff26dc0a7aa8d36e6cbf18ae1969c68c2_00.jpg&ehk=bsTR%2bc%2fHlXZOB2Odl3e3TrLATg2I2AVIeuviBlazrSo%3d&risl=&pid=ImgRaw&r=0",
    rating : 2
  },
  {
    autor : "- Portgas D Ace",
    quote : "Jangan pernah meremahkan diri sendiri, karena diri sendiri ada kelebihan tersendiri.",
    image : "https://th.bing.com/th/id/OIP.9B6_y5XVvKjG-L1NPNcJAAHaKe?pid=ImgDet&rs=1",
    rating : 3
  },
  {
    autor : "- Roronoa Zoro",
    quote : "Ketika dunia jahat kepadamu, maka berusahalah untuk menghadapinya, karena tidak ada orang yang membantumu jika kau tidak berusaha.",
    image : "https://i.pinimg.com/736x/bf/75/79/bf7579537070876642d7cc23e288fc2b.jpg",
    rating : 4
  },
  {
    autor : "- Nami cwannnn",
    quote : "Hidup ini seperti pensil yang lama lama akan habis, tetapi akan meninggalkan tulisan yang indah dalam kehidupan",
    image : "https://i.pinimg.com/736x/b2/80/da/b280da9df368c494eb6845cad594cac9.jpg",
    rating : 5
  }


]


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
// menampilkan defaultnya
allRating();

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
