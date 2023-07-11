class Rate {
    image = ""
    quote = ""
    name =""

    constructor(image,quote,name){
        this.image = image
        this.quote = quote
        this.name = name

    }
}
const test1 = new Rate('https://images.unsplash.com/photo-1686896435653-4a110ea92c84?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHx0b3BpYy1mZWVkfDEzfHRvd0paRnNrcEdnfHxlbnwwfHx8fHw%3D&auto=format&fit=crop&w=500&q=60','wah sangat keren yah gais','- fathul izhar')
const test2 = new Rate('https://images.unsplash.com/photo-1687444846017-1b4fd96a2743?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80','ini merupakan era dunia baru','- sasa irahi')
const test3 = new Rate('https://images.unsplash.com/photo-1686944716358-0fc13acb8266?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=938&q=80','dunia digital sangat membantu masyarakat berinteraksi lebih dalam','- oba chan')

let testimonial = [test1, test2, test3]
let testHTML = ''

for(let i = 0;i < testimonial.length; i++){
testHTML += `
<div class="card">
<img src="${testimonial[i].image}" alt="Avatar" style="width:100%">
<div class="container">
  <i>"${testimonial[i].quote}"</i> 
  <h4><b>${testimonial[i].name}</b></h4> 
</div>
</div>
`
}
document.getElementById('container-card').innerHTML = testHTML;
