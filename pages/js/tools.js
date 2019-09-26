function post(path, params, method='post') {
    const form = document.createElement('form');
    form.method = method;
    form.action = path;

    for (const key in params) {
        if (params.hasOwnProperty(key)) {
            const hiddenField = document.createElement('input');
            hiddenField.type = 'hidden';
            hiddenField.name = key;
            hiddenField.value = params[key];

            form.appendChild(hiddenField);
        }
    }

    document.body.appendChild(form);
    form.submit();
}

function setCookie(cookieName, cookieValue, nDays) {
    var today = new Date();
    var expire = new Date();

    if (nDays==null || nDays === 0)
        nDays=1;

    expire.setTime(today.getTime() + 3600000*24*nDays);
    document.cookie = cookieName+"="+escape(cookieValue) + ";expires="+expire.toGMTString();
}

function readCookie(name) {
    var nameEQ = name + "=";
    var ca = document.cookie.split(';');
    for(var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) === ' ') c = c.substring(1, c.length);
        if (c.indexOf(nameEQ) === 0) return c.substring(nameEQ.length, c.length);
    }
    return null;
}

let saveclass = null;
function saveTheme(cookieValue) {
    saveclass = saveclass ? saveclass : document.body.className;
    document.body.className = saveclass + ' ' + cookieValue;

    setCookie('theme', cookieValue, 365);
}

function onThemeChange() {
    let val = themeDropdown.dropdown('get value');
    editor.setTheme("ace/theme/" + val);
    saveTheme(val);
}

document.addEventListener('DOMContentLoaded', function() {
    let selectedTheme = readCookie('theme');

    themeDropdown.dropdown("set selected", selectedTheme);
    saveclass = saveclass ? saveclass : document.body.className;
    document.body.className = saveclass + ' ' + selectedTheme;

    editor.setTheme("ace/theme/" + themeDropdown.dropdown('get value'));
});