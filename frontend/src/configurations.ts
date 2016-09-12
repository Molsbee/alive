import * as ko from "knockout";
import * as $ from "jquery";

class ViewModel {
    configurations: KnockoutObservableArray<Configuration> = ko.observableArray<Configuration>();

    updateConfigurations(configs: Array<Configuration>): void {
        this.configurations(configs);
    }
}

interface Configuration {
    Id: number
    Endpoint: string
    ExpectedStatusCode: number
}

function populateConfigurations(model: any): void {
    console.log("Retrieving configurations");
    $.ajax({
        url: "/configuration/http",
        success: function(data) {
            let configs: Array<Configuration> = [];
            for (var i = 0; i < data.length; i++) {
                configs.push(data[i])
            }

            model.updateConfigurations(configs);
        }
    })
}

var viewModel = new ViewModel();
ko.applyBindings(viewModel);

populateConfigurations(viewModel);
