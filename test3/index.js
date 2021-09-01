//  font-familyを求める関数
function fontFamilyFunc() {
    let table = document.querySelector('table');
    let style = window.getComputedStyle(table);
    let fontFamily = style.getPropertyValue('font-family');

    return fontFamily;
}

//  font-sizeを求める関数
function fontSizeFunc() {
    let table = document.querySelector('table');
    let style = window.getComputedStyle(table).fontSize;
    let stFontSize = parseFloat(style);

    // console.log(stFontSize);

    return stFontSize;
}

//  フォントファミリーによってフォントサイズを変更する関数
function changeFontSize(fontFamily) {
    if (fontFamily == 'Caveat, Hannari') {
        document.querySelector('*').style.fontSize = '24px';
        document.querySelector('table input').style.fontSize = '24px';
    }
    if (fontFamily == '"Noto Sans JP", sans-serif') {
        document.querySelector('*').style.fontSize = '16px';
        document.querySelector('table input').style.fontSize = '16px';
    }
}

function main() {
    let fontFamily = fontFamilyFunc();

    let fontSize = fontSizeFunc();

    //  フォントファミリーによってフォントサイズを変更する関数
    changeFontSize(fontFamily);
}

main();
