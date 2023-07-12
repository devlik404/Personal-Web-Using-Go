class Rate {
    #quote = ""
    #image = ""
 
    constructor(quote,image){
       this.#quote = quote
        this.#image = image
       
    }
    get quote(){
      return this.#quote
    }
    get image(){
      return this.#image
    }
    get autor() {
      throw new Error('nama user harus di isi cokk')
    }
    get testHTML(){
      return `
      <div class="card">
      <img src="${this.image}" alt="Avatar" style="width:100%">
      <div class="container">
        <i>"${this.quote}"</i> 
        <h4><b>${this.autor}</b></h4> 
      </div>
      </div>
      `
    }
}
class testiRate extends Rate{
#autor = ""
constructor(autor,quote,image){
  super(quote,image)
  this.#autor = autor
}
get autor(){
  return "pengguna:" + this.#autor
}
}

class companyTest extends Rate{
  #company =""
  constructor(company,quote,image){
    super(quote,image)
    this.#company = company
  }

  get autor(){
    return "company:" + this.#company
  }
}
const test1 = new testiRate("- ikeh chan","kimochi desu","https://images.unsplash.com/photo-1686896435653-4a110ea92c84?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHx0b3BpYy1mZWVkfDEzfHRvd0paRnNrcEdnfHxlbnwwfHx8fHw%3D&auto=format&fit=crop&w=500&q=60")
const test2 = new testiRate('- miku',"sugggoi","https://images.unsplash.com/photo-1687444846017-1b4fd96a2743?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80")
const test3 = new companyTest("- pemuda pancasila","dimana ada rendang disitu ada kami","https://images.unsplash.com/photo-1686944716358-0fc13acb8266?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=938&q=80")


let testimonial = [test1, test2, test3]
let testHTML = ""

for(let i = 0;i < testimonial.length; i++){
testHTML += testimonial[i].testHTML
}
document.getElementById('container-card').innerHTML = testHTML;
