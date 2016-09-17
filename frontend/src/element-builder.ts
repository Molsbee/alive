export class Builder {

    element: HTMLElement;

    constructor(element: string) {
        this.element = document.createElement(element);
    }

    addClass(...clazz: string[]): Builder {
        this.element.classList.add(...clazz);
        return this;
    }

    setAttribute(name: string, value: string): Builder {
        this.element.setAttribute(name, value);
        return this;
    }

    appendChild(child: Node): Builder {
        this.element.appendChild(child);
        return this;
    }

    textContent(content: string): Builder {
        this.element.textContent = content;
        return this;
    }

    build(): HTMLElement {
        return this.element;
    }

}