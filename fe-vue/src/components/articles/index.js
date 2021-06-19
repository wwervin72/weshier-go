import { TweenMax, Expo, Elastic, Back, Power1 as Quad, Power3 as Quart } from 'gsap';

export const getOffset = (elem, axis) => {
	let offset = 0;
	const type = axis === "top" ? "offsetTop" : "offsetLeft";
	do {
		if (!isNaN(elem[type])) {
			offset += elem[type];
		}
	} while ((elem = elem.offsetParent));
	return offset;
};
export const distance = (p1, p2) => Math.hypot(p2.x - p1.x, p2.y - p1.y);
export const randNumber = (min, max) => Math.random() * (max - min + 1) + min;
export const getMousePos = e => {
	let posx = 0;
	let posy = 0;
	e = e || window.event;
	if (e.pageX || e.pageY) {
		posx = e.pageX;
		posy = e.pageY;
	} else if (e.clientX || e.clientY) {
		posx =
			e.clientX +
			document.body.scrollLeft +
			document.documentElement.scrollLeft;
		posy =
			e.clientY +
			document.body.scrollTop +
			document.documentElement.scrollTop;
	}
	return { x: posx, y: posy };
};
// Returns the rotation angle of an element.
export const getAngle = el => {
	const st = window.getComputedStyle(el, null);
	const tr = st.getPropertyValue("transform");
	let values = tr.split("(")[1];
	values = values.split(")")[0];
	values = values.split(",");
	return Math.round(Math.asin(values[1]) * (180 / Math.PI));
};
// Scroll control functions. Taken from https://stackoverflow.com/a/4770179.
const keys = { 37: 1, 38: 1, 39: 1, 40: 1 };
export const preventDefault = e => {
	e = e || window.event;
	if (e.preventDefault) e.preventDefault();
	e.returnValue = false;
};
export const preventDefaultForScrollKeys = e => {
	if (keys[e.keyCode]) {
		preventDefault(e);
		return false;
	}
};
export const disableScroll = () => {
	if (window.addEventListener)
		// older FF
		window.addEventListener("DOMMouseScroll", preventDefault, false);
	window.onwheel = preventDefault; // modern standard
	window.onmousewheel = document.onmousewheel = preventDefault; // older browsers, IE
	window.ontouchmove = preventDefault; // mobile
	document.onkeydown = preventDefaultForScrollKeys;
};
export const enableScroll = () => {
	if (window.removeEventListener)
		window.removeEventListener("DOMMouseScroll", preventDefault, false);
	window.onmousewheel = document.onmousewheel = null;
	window.onwheel = null;
	window.ontouchmove = null;
	document.onkeydown = null;
};
let allowTilt = true;
const body = document.body;
const docEl = document.documentElement;

let winsize;
const calcWinsize = () =>
	(winsize = { width: window.innerWidth, height: window.innerHeight });
calcWinsize();
window.addEventListener("resize", calcWinsize);

export class Content {
	constructor(el) {
		this.DOM = { el: el };
		this.DOM.img = this.DOM.el.querySelector(".content__item-img");
		this.DOM.title = this.DOM.el.querySelector(".content__item-title");
		this.DOM.text = this.DOM.el.querySelector(".content__item-text");
	}
	/**
	 * Show/Hide the content elements (title letters, the subtitle and the text).
	 */
	show(delay = 0, withAnimation = true) {
		this.toggle(delay, withAnimation);
	}
	hide(delay = 0, withAnimation = true) {
		this.toggle(delay, withAnimation, false);
	}
	toggle(delay, withAnimation, show = true) {
		setTimeout(
			() => {
				this.DOM.titleLetters.forEach((letter, pos) => {
					TweenMax.to(letter, !withAnimation ? 0 : show ? 0.6 : 0.3, {
						ease: show ? Back.easeOut : Quart.easeIn,
						delay: !withAnimation
							? 0
							: show
							? pos * 0.05
							: (this.titleLettersTotal - pos - 1) * 0.04,
						startAt: show ? { y: "50%", opacity: 0 } : null,
						y: show ? "0%" : "50%",
						opacity: show ? 1 : 0
					});
				});
				this.DOM.text.style.opacity = show ? 1 : 0;
			},
			withAnimation ? delay * 1000 : 0
		);
	}
}

class GridItem {
	constructor(el) {
		this.DOM = { el: el };
		this.DOM.bg = this.DOM.el.querySelector(".grid__item-bg");
		this.DOM.tilt = {};
		this.DOM.imgWrap = this.DOM.el.querySelector(".grid__item-wrap");
		this.DOM.tilt.img = this.DOM.imgWrap.querySelector("img");
		this.DOM.tilt.title = this.DOM.el.querySelector(".grid__item-title");
		this.DOM.tilt.content = this.DOM.el.querySelector(".grid__item-content");
		this.DOM.tilt.tag = this.DOM.el.querySelector(".grid__item-tag");
		this.DOM.tilt.info = this.DOM.el.querySelector(".grid__item-info");
		this.DOM.tags = this.DOM.tilt.tag.querySelectorAll("a");
		// 随机设置旋转角度
		const factor = Math.floor(Math.random() * (1 + 1)) || -1;
		this.DOM.tilt.img.style.cssText += `transform: rotate3d(0, 0, 1, ${factor *
			randNumber(1, 2)}deg);`;
		this.tiltconfig = {
			title: { translation: { x: [-8, 8], y: [4, -4] } },
			content: { translation: { x: [-10, 10], y: [6, -6] } },
			tag: { translation: { x: [-5, 5], y: [-10, 10] } },
			info: { translation: { x: [-5, 5], y: [-10, 10] } },
			img: { translation: { x: [-9, 9], y: [5, -5] } }
		};
		this.angle = getAngle(this.DOM.tilt.img);
		this.initEvents();
	}
	initEvents() {
		this.toggleAnimationOnHover = type => {
			TweenMax.to(this.DOM.bg, 1, {
				ease: Expo.easeOut,
				scale: type === "mouseenter" ? 1.15 : 1
			});
			this.DOM.tags.forEach((letter, pos) => {
				TweenMax.to(letter, 0.2, {
					ease: Quad.easeIn,
					delay: pos * 0.1,
					y: type === "mouseenter" ? "-50%" : "50%",
					opacity: 0,
					onComplete: () => {
						TweenMax.to(letter, type === "mouseenter" ? 0.6 : 1, {
							ease:
								type === "mouseenter"
									? Expo.easeOut
									: Elastic.easeOut.config(1, 0.4),
							startAt: {
								y: type === "mouseenter" ? "70%" : "-70%",
								opacity: 0
							},
							y: "0%",
							opacity: 1
						});
					}
				});
			});
		};
		this.mouseenterFn = ev => {
			if (!allowTilt) return;
			this.toggleAnimationOnHover(ev.type);
		};
		this.mousemoveFn = ev =>
			requestAnimationFrame(() => {
				if (!allowTilt) return;
				this.tilt(ev);
			});
		this.mouseleaveFn = ev => {
			if (!allowTilt) return;
			this.resetTilt();
			this.toggleAnimationOnHover(ev.type);
		};
		this.DOM.el.addEventListener("mouseenter", this.mouseenterFn);
		this.DOM.el.addEventListener("mousemove", this.mousemoveFn);
		this.DOM.el.addEventListener("mouseleave", this.mouseleaveFn);
	}
	tilt(ev) {
		const mousepos = getMousePos(ev);
		const docScrolls = {
			left: body.scrollLeft + docEl.scrollLeft,
			top: body.scrollTop + docEl.scrollTop
		};
		const bounds = this.DOM.el.getBoundingClientRect();
		const relmousepos = {
			x: mousepos.x - bounds.left - docScrolls.left,
			y: mousepos.y - bounds.top - docScrolls.top
		};
		for (let key in this.DOM.tilt) {
			let t = this.tiltconfig[key].translation;
			TweenMax.to(this.DOM.tilt[key], 2, {
				ease: Expo.easeOut,
				x: ((t.x[1] - t.x[0]) / bounds.width) * relmousepos.x + t.x[0],
				y: ((t.y[1] - t.y[0]) / bounds.height) * relmousepos.y + t.y[0]
			});
		}
	}
	resetTilt() {
		for (let key in this.DOM.tilt) {
			TweenMax.to(this.DOM.tilt[key], 2, {
				ease: Elastic.easeOut.config(1, 0.4),
				x: 0,
				y: 0
			});
		}
	}
	/**
	 * Hides the item:
	 * - Scales down and fades out the image and bg elements.
	 * - Moves down and fades out the title and number elements.
	 */
	hide(withAnimation = true) {
		this.toggle(withAnimation, false);
	}
	/**
	 * Resets.
	 */
	show(withAnimation = true) {
		this.toggle(withAnimation);
	}
	toggle(withAnimation, show = true) {
		TweenMax.to(this.DOM.tilt.img, withAnimation ? 0.8 : 0, {
			ease: Expo.easeInOut,
			delay: !withAnimation ? 0 : show ? 0.15 : 0,
			scale: show ? 1 : 0,
			opacity: show ? 1 : 0
		});
		TweenMax.to(this.DOM.bg, withAnimation ? 0.8 : 0, {
			ease: Expo.easeInOut,
			delay: !withAnimation ? 0 : show ? 0 : 0.15,
			scale: show ? 1 : 0,
			opacity: show ? 1 : 0
		});
		this.toggleItems(show ? 0.45 : 0, withAnimation, show);
	}
	// hides the texts (translate down and fade out).
	hideTexts(delay = 0, withAnimation = true) {
		this.toggleItems(delay, withAnimation, false);
	}
	// shows the texts (reset transforms and fade in).
	showTexts(delay = 0, withAnimation = true) {
		this.toggleItems(delay, withAnimation);
	}
	toggleItems(delay, withAnimation, show = true) {
		TweenMax.to(
			[this.DOM.tilt.title, this.DOM.tilt.tag, this.DOM.tilt.info, this.DOM.tilt.content],
			!withAnimation ? 0 : show ? 1 : 0.5,
			{
				ease: show ? Expo.easeOut : Quart.easeIn,
				delay: !withAnimation ? 0 : delay,
				y: show ? 0 : 20,
				opacity: show ? 1 : 0
			}
		);
	}
}

export class Grid {
	constructor(el, contentEl) {
		this.DOM = { el: el };
		this.contentEl = contentEl
		this.DOM.gridWrap = this.DOM.el.parentNode;
		this.items = [];
		Array.from(
			this.DOM.el.querySelectorAll(".grid__item")
		).forEach(itemEl => this.items.push(new GridItem(itemEl)));
		this.itemsTotal = this.items.length;
		this.DOM.closeCtrl = document.querySelector(".content__close");
		this.DOM.scrollIndicator = document.querySelector(
			".content__indicator"
		);
		this.current = -1;
	}
	initEvents(contentEl) {
		this.resizeFn = () => {
			if (this.current === -1 || winsize.width === window.innerWidth)
				return;
			this.closeItem(contentEl, false);
		};
		window.addEventListener("resize", this.resizeFn);
	}
	getSizePosition(el, scrolls = true) {
		const scrollTop =
			window.pageYOffset || docEl.scrollTop || body.scrollTop;
		const scrollLeft =
			window.pageXOffset || docEl.scrollLeft || body.scrollLeft;
		return {
			width: el.offsetWidth,
			height: el.offsetHeight,
			left: scrolls
				? getOffset(el, "left") - scrollLeft
				: getOffset(el, "left"),
			top: scrolls
				? getOffset(el, "top") - scrollTop
				: getOffset(el, "top")
		};
	}
	openItem(item, contentEl) {
		if (this.isAnimating) return;
		this.isAnimating = true;
		this.scrollPos = window.scrollY || window.pageYOffset || document.body.scrollTop;
		disableScroll();
		allowTilt = false;
		this.current = this.items.indexOf(item);
		this.hideAllItems(item);
		item.hideTexts();
		item.DOM.el.style.zIndex = 1000;
		const itemDim = this.getSizePosition(item.DOM.el);
		item.DOM.bg.style.width = `${itemDim.width}px`;
		item.DOM.bg.style.height = `${itemDim.height}px`;
		item.DOM.bg.style.left = `${itemDim.left}px`;
		item.DOM.bg.style.top = `${itemDim.top}px`;
		item.DOM.bg.style.position = "fixed";
		const d = Math.hypot(winsize.width, winsize.height);
		TweenMax.to(item.DOM.bg, 1.2, {
			ease: Expo.easeInOut,
			delay: 0.4,
			x: winsize.width / 2 - (itemDim.left + itemDim.width / 2),
			y: winsize.height / 2 - (itemDim.top + itemDim.height / 2),
			scaleX: d / itemDim.width,
			scaleY: d / itemDim.height,
			rotation: -1 * item.angle * 2
		});
		contentEl.DOM.el.classList.add("content__item--current");
		this.showContentElems(contentEl, 1);
		const imgDim = this.getSizePosition(item.DOM.imgWrap);
		const contentImgDim = this.getSizePosition(contentEl.DOM.img, false);
		TweenMax.to(item.DOM.tilt.img, 1.2, {
			ease: Expo.easeInOut,
			delay: 0.55,
			scaleX: contentImgDim.width / imgDim.width,
			scaleY: contentImgDim.height / imgDim.height,
			x:
				contentImgDim.left +
				contentImgDim.width / 2 -
				(imgDim.left + imgDim.width / 2),
			y:
				contentImgDim.top +
				contentImgDim.height / 2 -
				(imgDim.top + imgDim.height / 2),
			rotation: 0,
			onComplete: () => {
				item.DOM.tilt.img.style.opacity = 0;
				contentEl.DOM.img.style.visibility = "visible";
				contentEl.DOM.el.parentNode.style.position = "absolute";
				this.DOM.gridWrap.classList.add("grid-wrap--hidden");
				this.DOM.gridWrap.parentNode.classList.add('full_grid');
				window.scrollTo(0, 0);
				enableScroll();
				this.isAnimating = false;
			}
		});
	}
	closeItem(contentEl, withAnimation = true) {
		if (this.isAnimating) return;
		this.DOM.gridWrap.parentNode.classList.remove('full_grid');
		window.scrollTo(0, this.scrollPos);
		this.isAnimating = true;
		contentEl.DOM.el.parentNode.style.position = "fixed";
		disableScroll();
		this.DOM.gridWrap.classList.remove("grid-wrap--hidden");
		const item = this.items[this.current];
		this.hideContentElems(contentEl, 0, withAnimation);
		item.DOM.tilt.img.style.opacity = 1;
		contentEl.DOM.img.style.visibility = "hidden";
		TweenMax.to(item.DOM.tilt.img, withAnimation ? 1.2 : 0, {
			ease: Expo.easeInOut,
			scaleX: 1,
			scaleY: 1,
			x: 0,
			y: 0,
			rotation: item.angle * 2
		});
		TweenMax.to(item.DOM.bg, withAnimation ? 1.2 : 0, {
			ease: Expo.easeInOut,
			delay: 0.15,
			x: 0,
			y: 0,
			scaleX: 1,
			scaleY: 1,
			rotation: 0,
			onComplete: () => {
				contentEl.DOM.el.classList.remove("content__item--current");
				item.DOM.bg.style.position = "absolute";
				item.DOM.bg.style.left = "0px";
				item.DOM.bg.style.top = "0px";
				this.current = -1;
				allowTilt = true;
				item.DOM.el.style.zIndex = 0;
				enableScroll();
				this.isAnimating = false;
			}
		});
		this.showAllItems(item, withAnimation);
		item.showTexts(1, withAnimation);
	}
	/**
	 * Toggle the content elements.
	 */
	showContentElems(contentEl, delay = 0, withAnimation = true) {
		this.toggleContentElems(contentEl, delay, withAnimation);
	}
	hideContentElems(contentEl, delay = 0, withAnimation = true) {
		this.toggleContentElems(contentEl, delay, withAnimation, false);
	}
	toggleContentElems(contentEl, delay, withAnimation, show = true) {
		TweenMax.to(
			[this.DOM.closeCtrl, this.DOM.scrollIndicator],
			withAnimation ? 0.8 : 0,
			{
				ease: show ? Expo.easeOut : Expo.easeIn,
				delay: withAnimation ? delay : 0,
				startAt: show ? { y: 60 } : null,
				y: show ? 0 : 60,
				opacity: show ? 1 : 0
			}
		);
		if (show) {
			contentEl.show(delay, withAnimation);
		} else {
			contentEl.hide(delay, withAnimation);
		}
	}
	// Based on https://stackoverflow.com/q/25481717
	sortByDist(refPoint, itemsArray) {
		let distancePairs = [];
		let output = [];

		for (let i in itemsArray) {
			const rect = itemsArray[i].DOM.el.getBoundingClientRect();
			distancePairs.push([
				distance(refPoint, {
					x: rect.left + rect.width / 2,
					y: rect.top + rect.height / 2
				}),
				i
			]);
		}

		distancePairs.sort((a, b) => a[0] - b[0]);

		for (let p in distancePairs) {
			const pair = distancePairs[p];
			output.push(itemsArray[pair[1]]);
		}

		return output;
	}
	/**
	 * Shows/Hides all the grid items except the "exclude" item.
	 * The items will be sorted based on the distance to the exclude item.
	 */
	showAllItems(exclude, withAnimation = true) {
		this.toggleAllItems(exclude, withAnimation);
	}
	hideAllItems(exclude, withAnimation = true) {
		this.toggleAllItems(exclude, withAnimation, false);
	}
	toggleAllItems(exclude, withAnimation, show = true) {
		if (!withAnimation) {
			this.items
				.filter(item => item != exclude)
				.forEach(item => item[show ? "show" : "hide"](withAnimation));
		} else {
			const refrect = exclude.DOM.el.getBoundingClientRect();
			const refPoint = {
				x: refrect.left + refrect.width / 2,
				y: refrect.top + refrect.height / 2
			};
			this.sortByDist(
				refPoint,
				this.items.filter(item => item != exclude)
			).forEach((item, pos) =>
				setTimeout(
					() => item[show ? "show" : "hide"](),
					show ? 300 + pos * 70 : pos * 70
				)
			);
		}
	}
}
