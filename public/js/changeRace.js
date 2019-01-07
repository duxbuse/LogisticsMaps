function changeRace() {
    var x = document.getElementsByName("FRace");
    var FriendlyRace = x[0][x[0].selectedIndex].text

    list = document.getElementsByClassName("FSpecialties")
    for (let element of list) {
        

        //true if element is of the current race or "any"
        var b = element.classList.contains(FriendlyRace) || element.classList.contains("any");

        if(b){
            element.style.display = 'block'
        }else{
            element.style.display = 'none'
        }
    }

    //repeated for Enemy
    var x = document.getElementsByName("ERace");
    var EnemyRace = x[0][x[0].selectedIndex].text

    list = document.getElementsByClassName("ESpecialties")
    for (let element of list) {
        

        //true if element is of the current race or "any"
        var b = element.classList.contains(EnemyRace) || element.classList.contains("any");

        if(b){
            element.style.display = 'block'
        }else{
            element.style.display = 'none'
        }
    }
}