window.addEventListener("load", () => {
	let editor = ace.edit("editor");
	editor.setTheme("ace/theme/monokai");
	editor.setFontSize(14);
	editor.getSession().setMode("ace/mode/html");
	editor.getSession().setUseWrapMode(true); //折り返し
	editor.getSession().setTabSize(2); //タブ幅
})

