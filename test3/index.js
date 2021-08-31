function main() {
    let table = document.querySelector('table');
    let style = window.getComputedStyle(table)
    let fontFamily = style.getPropertyValue("font-family")
    if (fontFamily == "Noto Sans JP")
}

main();
