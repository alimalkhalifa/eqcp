// source: proto/petition.proto
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

goog.exportSymbol('proto.pb.Petition', null, global);
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
proto.pb.Petition = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.Petition, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.Petition.displayName = 'proto.pb.Petition';
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
proto.pb.Petition.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.Petition.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.Petition} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.Petition.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    petid: jspb.Message.getFieldWithDefault(msg, 2, 0),
    charname: jspb.Message.getFieldWithDefault(msg, 3, ""),
    accountname: jspb.Message.getFieldWithDefault(msg, 4, ""),
    lastgm: jspb.Message.getFieldWithDefault(msg, 5, ""),
    petitiontext: jspb.Message.getFieldWithDefault(msg, 6, ""),
    gmtext: jspb.Message.getFieldWithDefault(msg, 7, ""),
    zone: jspb.Message.getFieldWithDefault(msg, 8, ""),
    urgency: jspb.Message.getFieldWithDefault(msg, 9, 0),
    charclass: jspb.Message.getFieldWithDefault(msg, 10, 0),
    charrace: jspb.Message.getFieldWithDefault(msg, 11, 0),
    charlevel: jspb.Message.getFieldWithDefault(msg, 12, 0),
    checkouts: jspb.Message.getFieldWithDefault(msg, 13, 0),
    unavailables: jspb.Message.getFieldWithDefault(msg, 14, 0),
    ischeckedout: jspb.Message.getFieldWithDefault(msg, 15, 0),
    senttime: jspb.Message.getFieldWithDefault(msg, 16, 0)
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
 * @return {!proto.pb.Petition}
 */
proto.pb.Petition.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.Petition;
  return proto.pb.Petition.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.Petition} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.Petition}
 */
proto.pb.Petition.deserializeBinaryFromReader = function(msg, reader) {
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
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPetid(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setCharname(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountname(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastgm(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setPetitiontext(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setGmtext(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setZone(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUrgency(value);
      break;
    case 10:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCharclass(value);
      break;
    case 11:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCharrace(value);
      break;
    case 12:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCharlevel(value);
      break;
    case 13:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCheckouts(value);
      break;
    case 14:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUnavailables(value);
      break;
    case 15:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setIscheckedout(value);
      break;
    case 16:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setSenttime(value);
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
proto.pb.Petition.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.Petition.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.Petition} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.Petition.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getPetid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getCharname();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getAccountname();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLastgm();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getPetitiontext();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getGmtext();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getZone();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getUrgency();
  if (f !== 0) {
    writer.writeInt64(
      9,
      f
    );
  }
  f = message.getCharclass();
  if (f !== 0) {
    writer.writeInt64(
      10,
      f
    );
  }
  f = message.getCharrace();
  if (f !== 0) {
    writer.writeInt64(
      11,
      f
    );
  }
  f = message.getCharlevel();
  if (f !== 0) {
    writer.writeInt64(
      12,
      f
    );
  }
  f = message.getCheckouts();
  if (f !== 0) {
    writer.writeInt64(
      13,
      f
    );
  }
  f = message.getUnavailables();
  if (f !== 0) {
    writer.writeInt64(
      14,
      f
    );
  }
  f = message.getIscheckedout();
  if (f !== 0) {
    writer.writeInt64(
      15,
      f
    );
  }
  f = message.getSenttime();
  if (f !== 0) {
    writer.writeInt64(
      16,
      f
    );
  }
};


/**
 * optional int64 id = 1;
 * @return {number}
 */
proto.pb.Petition.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 petid = 2;
 * @return {number}
 */
proto.pb.Petition.prototype.getPetid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setPetid = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional string charname = 3;
 * @return {string}
 */
proto.pb.Petition.prototype.getCharname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setCharname = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string accountname = 4;
 * @return {string}
 */
proto.pb.Petition.prototype.getAccountname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setAccountname = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string lastgm = 5;
 * @return {string}
 */
proto.pb.Petition.prototype.getLastgm = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setLastgm = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string petitiontext = 6;
 * @return {string}
 */
proto.pb.Petition.prototype.getPetitiontext = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setPetitiontext = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string gmtext = 7;
 * @return {string}
 */
proto.pb.Petition.prototype.getGmtext = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setGmtext = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string zone = 8;
 * @return {string}
 */
proto.pb.Petition.prototype.getZone = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setZone = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional int64 urgency = 9;
 * @return {number}
 */
proto.pb.Petition.prototype.getUrgency = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setUrgency = function(value) {
  return jspb.Message.setProto3IntField(this, 9, value);
};


/**
 * optional int64 charclass = 10;
 * @return {number}
 */
proto.pb.Petition.prototype.getCharclass = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 10, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setCharclass = function(value) {
  return jspb.Message.setProto3IntField(this, 10, value);
};


/**
 * optional int64 charrace = 11;
 * @return {number}
 */
proto.pb.Petition.prototype.getCharrace = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 11, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setCharrace = function(value) {
  return jspb.Message.setProto3IntField(this, 11, value);
};


/**
 * optional int64 charlevel = 12;
 * @return {number}
 */
proto.pb.Petition.prototype.getCharlevel = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 12, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setCharlevel = function(value) {
  return jspb.Message.setProto3IntField(this, 12, value);
};


/**
 * optional int64 checkouts = 13;
 * @return {number}
 */
proto.pb.Petition.prototype.getCheckouts = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 13, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setCheckouts = function(value) {
  return jspb.Message.setProto3IntField(this, 13, value);
};


/**
 * optional int64 unavailables = 14;
 * @return {number}
 */
proto.pb.Petition.prototype.getUnavailables = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 14, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setUnavailables = function(value) {
  return jspb.Message.setProto3IntField(this, 14, value);
};


/**
 * optional int64 ischeckedout = 15;
 * @return {number}
 */
proto.pb.Petition.prototype.getIscheckedout = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 15, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setIscheckedout = function(value) {
  return jspb.Message.setProto3IntField(this, 15, value);
};


/**
 * optional int64 senttime = 16;
 * @return {number}
 */
proto.pb.Petition.prototype.getSenttime = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 16, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.Petition} returns this
 */
proto.pb.Petition.prototype.setSenttime = function(value) {
  return jspb.Message.setProto3IntField(this, 16, value);
};


goog.object.extend(exports, proto.pb);
