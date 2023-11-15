const button = document.querySelector(".siycb3m button");
button != undefined ? button.click() : null; 
const termArr = document.querySelectorAll('.SetPageTerm-wordText span');
const defArr = document.querySelectorAll('.SetPageTerm-definitionText span');
let out = "[\n";
termArr.forEach((elm, index) => {
    out += `\t "${elm.innerText}" = "${defArr[index].innerText}", \n`
})
out += "]";
console.log(out);