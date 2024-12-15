import { createIcons, icons } from "lucide";

$(() => {
	createIcons({icons});
});


// Update favicon
const iconLight = document.querySelector("link#favicon-light");
const iconDark = document.querySelector("link#favicon-dark");

$(() => {
	var matcher = window.matchMedia("(prefers-color-scheme: dark)");
	matcher.addListener(updateIcon);
	updateIcon();

	function updateIcon() {
		if (matcher.matches) {
			iconLight.remove();
			document.head.append(iconDark);
		} else {
			document.head.append(iconLight);
			iconDark.remove();
		}
	}
});
