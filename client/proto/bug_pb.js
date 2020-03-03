// source: proto/bug.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.pb.Bug', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.Bug = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.Bug, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.Bug.displayName = 'proto.pb.Bug';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.Bug.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.Bug.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.Bug} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.Bug.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    zone: jspb.Message.getFieldWithDefault(msg, 2, ""),
    clientversionid: jspb.Message.getFieldWithDefault(msg, 3, 0),
    clientversionname: jspb.Message.getFieldWithDefault(msg, 4, ""),
    accountid: jspb.Message.getFieldWithDefault(msg, 5, 0),
    characterid: jspb.Message.getFieldWithDefault(msg, 6, 0),
    charactername: jspb.Message.getFieldWithDefault(msg, 7, ""),
    reporterspoof: jspb.Message.getFieldWithDefault(msg, 8, 0),
    categoryid: jspb.Message.getFieldWithDefault(msg, 9, 0),
    categoryname: jspb.Message.getFieldWithDefault(msg, 10, ""),
    reportername: jspb.Message.getFieldWithDefault(msg, 11, ""),
    uipath: jspb.Message.getFieldWithDefault(msg, 12, ""),
    posx: jspb.Message.getFieldWithDefault(msg, 13, 0),
    posy: jspb.Message.getFieldWithDefault(msg, 14, 0),
    posz: jspb.Message.getFieldWithDefault(msg, 15, 0),
    heading: jspb.Message.getFieldWithDefault(msg, 16, 0),
    timeplayed: jspb.Message.getFieldWithDefault(msg, 17, 0),
    targetid: jspb.Message.getFieldWithDefault(msg, 18, 0),
    targetname: jspb.Message.getFieldWithDefault(msg, 19, ""),
    optionalinfomask: jspb.Message.getFieldWithDefault(msg, 20, 0),
    canduplicate: jspb.Message.getFieldWithDefault(msg, 21, 0),
    crashbug: jspb.Message.getFieldWithDefault(msg, 22, 0),
    targetinfo: jspb.Message.getFieldWithDefault(msg, 23, 0),
    characterflags: jspb.Message.getFieldWithDefault(msg, 24, 0),
    unknownvalue: jspb.Message.getFieldWithDefault(msg, 25, 0),
    bugreport: jspb.Message.getFieldWithDefault(msg, 26, ""),
    systeminfo: jspb.Message.getFieldWithDefault(msg, 27, ""),
    reportdatetime: jspb.Message.getFieldWithDefault(msg, 28, 0),
    bugstatus: jspb.Message.getFieldWithDefault(msg, 29, 0),
    lastreview: jspb.Message.getFieldWithDefault(msg, 30, 0),
    lastreviewer: jspb.Message.getFieldWithDefault(msg, 31, ""),
    reviewernotes: jspb.Message.getFieldWithDefault(msg, 32, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.Bug}
 */
proto.pb.Bug.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.Bug;
  return proto.pb.Bug.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.Bug} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.Bug}
 */
proto.pb.Bug.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setZone(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setClientversionid(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setClientversionname(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setAccountid(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCharacterid(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setCharactername(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setReporterspoof(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCategoryid(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setCategoryname(value);
      break;
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setReportername(value);
      break;
    case 12:
      var value = /** @type {string} */ (reader.readString());
      msg.setUipath(value);
      break;
    case 13:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPosx(value);
      break;
    case 14:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPosy(value);
      break;
    case 15:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPosz(value);
      break;
    case 16:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setHeading(value);
      break;
    case 17:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTimeplayed(value);
      break;
    case 18:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTargetid(value);
      break;
    case 19:
      var value = /** @type {string} */ (reader.readString());
      msg.setTargetname(value);
      break;
    case 20:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setOptionalinfomask(value);
      break;
    case 21:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCanduplicate(value);
      break;
    case 22:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCrashbug(value);
      break;
    case 23:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTargetinfo(value);
      break;
    case 24:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCharacterflags(value);
      break;
    case 25:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUnknownvalue(value);
      break;
    case 26:
      var value = /** @type {string} */ (reader.readString());
      msg.setBugreport(value);
      break;
    case 27:
      var value = /** @type {string} */ (reader.readString());
      msg.setSysteminfo(value);
      break;
    case 28:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setReportdatetime(value);
      break;
    case 29:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setBugstatus(value);
      break;
    case 30:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setLastreview(value);
      break;
    case 31:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastreviewer(value);
      break;
    case 32:
      var value = /** @type {string} */ (reader.readString());
      msg.setReviewernotes(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.Bug.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.Bug.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.Bug} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.Bug.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getZone();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getClientversionid();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getClientversionname();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getAccountid();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
  f = message.getCharacterid();
  if (f !== 0) {
    writer.writeInt64(
      6,
      f
    );
  }
  f = message.getCharactername();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getReporterspoof();
  if (f !== 0) {
    writer.writeInt64(
      8,
      f
    );
  }
  f = message.getCategoryid();
  if (f !== 0) {
    writer.writeInt64(
      9,
      f
    );
  }
  f = message.getCategoryname();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
  f = message.getReportername();
  if (f.length > 0) {
    writer.writeString(
      11,
      f
    );
  }
  f = message.getUipath();
  if (f.length > 0) {
    writer.writeString(
      12,
      f
    );
  }
  f = message.getPosx();
  if (f !== 0) {
    writer.writeInt64(
      13,
      f
    );
  }
  f = message.getPosy();
  if (f !== 0) {
    writer.writeInt64(
      14,
      f
    );
  }
  f = message.getPosz();
  if (f !== 0) {
    writer.writeInt64(
      15,
      f
    );
  }
  f = message.getHeading();
  if (f !== 0) {
    writer.writeInt64(
      16,
      f
    );
  }
  f = message.getTimeplayed();
  if (f !== 0) {
    writer.writeInt64(
      17,
      f
    );
  }
  f = message.getTargetid();
  if (f !== 0) {
    writer.writeInt64(
      18,
      f
    );
  }
  f = message.getTargetname();
  if (f.length > 0) {
    writer.writeString(
      19,
      f
    );
  }
  f = message.getOptionalinfomask();
  if (f !== 0) {
    writer.writeInt64(
      20,
      f
    );
  }
  f = message.getCanduplicate();
  if (f !== 0) {
    writer.writeInt64(
      21,
      f
    );
  }
  f = message.getCrashbug();
  if (f !== 0) {
    writer.writeInt64(
      22,
      f
    );
  }
  f = message.getTargetinfo();
  if (f !== 0) {
    writer.writeInt64(
      23,
      f
    );
  }
  f = message.getCharacterflags();
  if (f !== 0) {
    writer.writeInt64(
      24,
      f
    );
  }
  f = message.getUnknownvalue();
  if (f !== 0) {
    writer.writeInt64(
      25,
      f
    );
  }
  f = message.getBugreport();
  if (f.length > 0) {
    writer.writeString(
      26,
      f
    );
  }
  f = message.getSysteminfo();
  if (f.length > 0) {
    writer.writeString(
      27,
      f
    );
  }
  f = message.getReportdatetime();
  if (f !== 0) {
    writer.writeInt64(
      28,
      f
    );
  }
  f = message.getBugstatus();
  if (f !== 0) {
    writer.writeInt64(
      29,
      f
    );
  }
  f = message.getLastreview();
  if (f !== 0) {
    writer.writeInt64(
      30,
      f
    );
  }
  f = message.getLastreviewer();
  if (f.length > 0) {
    writer.writeString(
      31,
      f
    );
  }
  f = message.getReviewernotes();
  if (f.length > 0) {
    writer.writeString(
      32,
      f
    );
  }
};


/**
 * optional int64 id = 1;
 * @return {number}
 */
proto.pb.Bug.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string zone = 2;
 * @return {string}
 */
proto.pb.Bug.prototype.getZone = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setZone = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int64 clientversionid = 3;
 * @return {number}
 */
proto.pb.Bug.prototype.getClientversionid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setClientversionid = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional string clientversionname = 4;
 * @return {string}
 */
proto.pb.Bug.prototype.getClientversionname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setClientversionname = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional int64 accountid = 5;
 * @return {number}
 */
proto.pb.Bug.prototype.getAccountid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setAccountid = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional int64 characterid = 6;
 * @return {number}
 */
proto.pb.Bug.prototype.getCharacterid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setCharacterid = function(value) {
  return jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional string charactername = 7;
 * @return {string}
 */
proto.pb.Bug.prototype.getCharactername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setCharactername = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional int64 reporterspoof = 8;
 * @return {number}
 */
proto.pb.Bug.prototype.getReporterspoof = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setReporterspoof = function(value) {
  return jspb.Message.setProto3IntField(this, 8, value);
};


/**
 * optional int64 categoryid = 9;
 * @return {number}
 */
proto.pb.Bug.prototype.getCategoryid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setCategoryid = function(value) {
  return jspb.Message.setProto3IntField(this, 9, value);
};


/**
 * optional string categoryname = 10;
 * @return {string}
 */
proto.pb.Bug.prototype.getCategoryname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setCategoryname = function(value) {
  return jspb.Message.setProto3StringField(this, 10, value);
};


/**
 * optional string reportername = 11;
 * @return {string}
 */
proto.pb.Bug.prototype.getReportername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setReportername = function(value) {
  return jspb.Message.setProto3StringField(this, 11, value);
};


/**
 * optional string uipath = 12;
 * @return {string}
 */
proto.pb.Bug.prototype.getUipath = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 12, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setUipath = function(value) {
  return jspb.Message.setProto3StringField(this, 12, value);
};


/**
 * optional int64 posx = 13;
 * @return {number}
 */
proto.pb.Bug.prototype.getPosx = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 13, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setPosx = function(value) {
  return jspb.Message.setProto3IntField(this, 13, value);
};


/**
 * optional int64 posy = 14;
 * @return {number}
 */
proto.pb.Bug.prototype.getPosy = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 14, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setPosy = function(value) {
  return jspb.Message.setProto3IntField(this, 14, value);
};


/**
 * optional int64 posz = 15;
 * @return {number}
 */
proto.pb.Bug.prototype.getPosz = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 15, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setPosz = function(value) {
  return jspb.Message.setProto3IntField(this, 15, value);
};


/**
 * optional int64 heading = 16;
 * @return {number}
 */
proto.pb.Bug.prototype.getHeading = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 16, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setHeading = function(value) {
  return jspb.Message.setProto3IntField(this, 16, value);
};


/**
 * optional int64 timeplayed = 17;
 * @return {number}
 */
proto.pb.Bug.prototype.getTimeplayed = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 17, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setTimeplayed = function(value) {
  return jspb.Message.setProto3IntField(this, 17, value);
};


/**
 * optional int64 targetid = 18;
 * @return {number}
 */
proto.pb.Bug.prototype.getTargetid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 18, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setTargetid = function(value) {
  return jspb.Message.setProto3IntField(this, 18, value);
};


/**
 * optional string targetname = 19;
 * @return {string}
 */
proto.pb.Bug.prototype.getTargetname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 19, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setTargetname = function(value) {
  return jspb.Message.setProto3StringField(this, 19, value);
};


/**
 * optional int64 optionalinfomask = 20;
 * @return {number}
 */
proto.pb.Bug.prototype.getOptionalinfomask = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 20, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setOptionalinfomask = function(value) {
  return jspb.Message.setProto3IntField(this, 20, value);
};


/**
 * optional int64 canduplicate = 21;
 * @return {number}
 */
proto.pb.Bug.prototype.getCanduplicate = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 21, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setCanduplicate = function(value) {
  return jspb.Message.setProto3IntField(this, 21, value);
};


/**
 * optional int64 crashbug = 22;
 * @return {number}
 */
proto.pb.Bug.prototype.getCrashbug = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 22, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setCrashbug = function(value) {
  return jspb.Message.setProto3IntField(this, 22, value);
};


/**
 * optional int64 targetinfo = 23;
 * @return {number}
 */
proto.pb.Bug.prototype.getTargetinfo = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 23, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setTargetinfo = function(value) {
  return jspb.Message.setProto3IntField(this, 23, value);
};


/**
 * optional int64 characterflags = 24;
 * @return {number}
 */
proto.pb.Bug.prototype.getCharacterflags = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 24, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setCharacterflags = function(value) {
  return jspb.Message.setProto3IntField(this, 24, value);
};


/**
 * optional int64 unknownvalue = 25;
 * @return {number}
 */
proto.pb.Bug.prototype.getUnknownvalue = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 25, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setUnknownvalue = function(value) {
  return jspb.Message.setProto3IntField(this, 25, value);
};


/**
 * optional string bugreport = 26;
 * @return {string}
 */
proto.pb.Bug.prototype.getBugreport = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 26, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setBugreport = function(value) {
  return jspb.Message.setProto3StringField(this, 26, value);
};


/**
 * optional string systeminfo = 27;
 * @return {string}
 */
proto.pb.Bug.prototype.getSysteminfo = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 27, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setSysteminfo = function(value) {
  return jspb.Message.setProto3StringField(this, 27, value);
};


/**
 * optional int64 reportdatetime = 28;
 * @return {number}
 */
proto.pb.Bug.prototype.getReportdatetime = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 28, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setReportdatetime = function(value) {
  return jspb.Message.setProto3IntField(this, 28, value);
};


/**
 * optional int64 bugstatus = 29;
 * @return {number}
 */
proto.pb.Bug.prototype.getBugstatus = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 29, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setBugstatus = function(value) {
  return jspb.Message.setProto3IntField(this, 29, value);
};


/**
 * optional int64 lastreview = 30;
 * @return {number}
 */
proto.pb.Bug.prototype.getLastreview = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 30, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setLastreview = function(value) {
  return jspb.Message.setProto3IntField(this, 30, value);
};


/**
 * optional string lastreviewer = 31;
 * @return {string}
 */
proto.pb.Bug.prototype.getLastreviewer = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 31, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setLastreviewer = function(value) {
  return jspb.Message.setProto3StringField(this, 31, value);
};


/**
 * optional string reviewernotes = 32;
 * @return {string}
 */
proto.pb.Bug.prototype.getReviewernotes = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 32, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Bug} returns this
 */
proto.pb.Bug.prototype.setReviewernotes = function(value) {
  return jspb.Message.setProto3StringField(this, 32, value);
};


goog.object.extend(exports, proto.pb);
