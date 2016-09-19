import * as element from "./element-builder";

class Navigation {

    nav: HTMLElement;
    sideBar: HTMLElement;

    constructor() {
        var iconBar = new element.Builder("span").addClass("icon-bar").build();

        this.nav = new element.Builder("nav")
            .addClass("navbar", "navbar-inverse", "navbar-fixed-top")
            .appendChild(
                new element.Builder("div")
                    .addClass("container-fluid")
                    .appendChild(
                        new element.Builder("div")
                            .addClass("navbar-header")
                            .appendChild(
                                new element.Builder("button")
                                    .addClass("navbar-toggle", "collapsed")
                                    .setAttribute("data-toggle", "collapse")
                                    .setAttribute("data-target", "#navbar")
                                    .setAttribute("aria-expanded", "false")
                                    .setAttribute("aria-controls", "navbar")
                                    .appendChild(
                                        new element.Builder("span")
                                            .addClass("sr-only")
                                            .textContent("Toggle navigation")
                                            .build()
                                    )
                                    .appendChild(
                                        new element.Builder("span")
                                            .addClass("icon-bar")
                                            .build()
                                    )
                                    .appendChild(
                                        new element.Builder("span")
                                            .addClass("icon-bar")
                                            .build()
                                    )
                                    .appendChild(
                                        new element.Builder("span")
                                            .addClass("icon-bar")
                                            .build()
                                    )
                                    .build()
                            )
                            .build()
                    )
                    .appendChild(
                        new element.Builder("a")
                            .addClass("navbar-brand")
                            .setAttribute("href", "/")
                            .textContent("Alive")
                            .build()
                    )
                    .build()
            )
            .appendChild(
                new element.Builder("div")
                    .addClass("navbar-collapse", "collapse")
                    .build()
            )
            .build();


        this.sideBar = new element.Builder("div")
            .appendChild(
                new element.Builder("ul")
                    .addClass("nav")
                    .appendChild(
                        createSideBarItem("glyphicon-home", "dashboard")
                    ).appendChild(
                        createSideBarItem("glyphicon-plus", "create")
                    ).appendChild(
                        createSideBarItem("glyphicon-list-alt", "monitoring")
                    ).appendChild(
                        createSideBarItem("glyphicon-cog", "settings")
                    )
                    .build()
            )
            .build();
    }

}

function createSideBarItem(glyphicon: string, text: string): HTMLElement {
    return new element.Builder("li")
                .appendChild(
                    new element.Builder("a")
                        .setAttribute("href", "#")
                        .appendChild(
                            new element.Builder("span")
                                .addClass("glyphicon", glyphicon)
                                .setAttribute("aria-hidden", "true")
                                .build()
                        )
                        .appendChild(
                            new element.Builder("span")
                                .textContent(" " + text)
                                .build()
                        )
                        .build()
                ).build()
}

var navigation = new Navigation();
var brandBar = document.getElementsByTagName("brand-bar")[0];
brandBar.appendChild(navigation.nav);

var sideBar = document.getElementsByTagName("side-bar")[0];
sideBar.appendChild(navigation.sideBar);

var main = document.getElementsByTagName("main")[0];
main.classList.add("col-md-offset-2", "container-fluid");
