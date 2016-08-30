var PingResponse = function(data) {
  var self = this;

  self.id = data.id;
  self.createdTS = moment.utc(data.createdTS).toDate();
  self.durationMS = data.durationMS;
  self.statusCode = data.statusCode;
  self.pingConfigId = data.pingConfigId;

  return self;
}
