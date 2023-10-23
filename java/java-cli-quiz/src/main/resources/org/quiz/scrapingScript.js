const button = document.querySelector(".siycb3m button");
button != undefined ? button.click() : null;
const termArr = document.querySelectorAll('.SetPageTerm-wordText span');
const defArr = document.querySelectorAll('.SetPageTerm-definitionText span');
let out = "";
termArr.forEach((elm, index) => {
    out += `${elm.innerText}=${defArr[index].innerText};`
})
console.log(out);