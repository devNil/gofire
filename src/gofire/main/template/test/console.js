function Output(id){
     this.out = document.getElementById(id);
}
            
Output.prototype.add = function(text){
    var art = document.createElement("article");
    art.innerText = text;
    this.out.appendChild(art);
}