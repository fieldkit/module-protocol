/* Automatically generated nanopb constant definitions */
/* Generated by nanopb-0.3.9-dev at Mon Dec  4 10:50:23 2017. */

#include "fk-module.pb.h"

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif



const pb_field_t fk_module_QueryCapabilities_fields[3] = {
    PB_FIELD(  1, INT32   , SINGULAR, STATIC  , FIRST, fk_module_QueryCapabilities, version, version, 0),
    PB_FIELD(  2, UINT32  , SINGULAR, STATIC  , OTHER, fk_module_QueryCapabilities, callerTime, version, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_module_QuerySensorCapabilities_fields[2] = {
    PB_FIELD(  1, INT32   , SINGULAR, STATIC  , FIRST, fk_module_QuerySensorCapabilities, sensor, sensor, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_module_Capabilities_fields[5] = {
    PB_FIELD(  1, INT32   , SINGULAR, STATIC  , FIRST, fk_module_Capabilities, version, version, 0),
    PB_FIELD(  2, UENUM   , SINGULAR, STATIC  , OTHER, fk_module_Capabilities, type, version, 0),
    PB_FIELD(  3, STRING  , SINGULAR, CALLBACK, OTHER, fk_module_Capabilities, name, type, 0),
    PB_FIELD(  4, INT32   , SINGULAR, STATIC  , OTHER, fk_module_Capabilities, numberOfSensors, name, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_module_SensorCapabilities_fields[4] = {
    PB_FIELD(  1, INT32   , SINGULAR, STATIC  , FIRST, fk_module_SensorCapabilities, id, id, 0),
    PB_FIELD(  2, STRING  , SINGULAR, CALLBACK, OTHER, fk_module_SensorCapabilities, name, id, 0),
    PB_FIELD(  3, STRING  , SINGULAR, CALLBACK, OTHER, fk_module_SensorCapabilities, unitOfMeasure, name, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_module_BeginTakeReadings_fields[2] = {
    PB_FIELD(  1, INT32   , SINGULAR, STATIC  , FIRST, fk_module_BeginTakeReadings, index, index, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_module_QueryReadingStatus_fields[1] = {
    PB_LAST_FIELD
};

const pb_field_t fk_module_ReadingStatus_fields[2] = {
    PB_FIELD(  1, UENUM   , SINGULAR, STATIC  , FIRST, fk_module_ReadingStatus, state, state, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_module_SensorReadings_fields[1] = {
    PB_LAST_FIELD
};

const pb_field_t fk_module_SensorReading_fields[4] = {
    PB_FIELD(  1, INT32   , SINGULAR, STATIC  , FIRST, fk_module_SensorReading, sensor, sensor, 0),
    PB_FIELD(  2, UINT32  , SINGULAR, STATIC  , OTHER, fk_module_SensorReading, time, sensor, 0),
    PB_FIELD(  3, FLOAT   , SINGULAR, STATIC  , OTHER, fk_module_SensorReading, value, time, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_module_WireMessageQuery_fields[6] = {
    PB_FIELD(  1, UENUM   , SINGULAR, STATIC  , FIRST, fk_module_WireMessageQuery, type, type, 0),
    PB_FIELD(  2, MESSAGE , SINGULAR, STATIC  , OTHER, fk_module_WireMessageQuery, queryCapabilities, type, &fk_module_QueryCapabilities_fields),
    PB_FIELD(  3, MESSAGE , SINGULAR, STATIC  , OTHER, fk_module_WireMessageQuery, querySensorCapabilities, queryCapabilities, &fk_module_QuerySensorCapabilities_fields),
    PB_FIELD(  4, MESSAGE , SINGULAR, STATIC  , OTHER, fk_module_WireMessageQuery, beginTakeReadings, querySensorCapabilities, &fk_module_BeginTakeReadings_fields),
    PB_FIELD(  5, MESSAGE , SINGULAR, STATIC  , OTHER, fk_module_WireMessageQuery, queryReadingStatus, beginTakeReadings, &fk_module_QueryReadingStatus_fields),
    PB_LAST_FIELD
};

const pb_field_t fk_module_Error_fields[3] = {
    PB_FIELD(  1, UENUM   , SINGULAR, STATIC  , FIRST, fk_module_Error, type, type, 0),
    PB_FIELD(  2, STRING  , SINGULAR, CALLBACK, OTHER, fk_module_Error, message, type, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_module_WireMessageReply_fields[7] = {
    PB_FIELD(  1, UENUM   , SINGULAR, STATIC  , FIRST, fk_module_WireMessageReply, type, type, 0),
    PB_FIELD(  2, MESSAGE , SINGULAR, STATIC  , OTHER, fk_module_WireMessageReply, error, type, &fk_module_Error_fields),
    PB_FIELD(  3, MESSAGE , SINGULAR, STATIC  , OTHER, fk_module_WireMessageReply, capabilities, error, &fk_module_Capabilities_fields),
    PB_FIELD(  4, MESSAGE , SINGULAR, STATIC  , OTHER, fk_module_WireMessageReply, sensorCapabilities, capabilities, &fk_module_SensorCapabilities_fields),
    PB_FIELD(  5, MESSAGE , SINGULAR, STATIC  , OTHER, fk_module_WireMessageReply, readingStatus, sensorCapabilities, &fk_module_ReadingStatus_fields),
    PB_FIELD(  6, MESSAGE , SINGULAR, STATIC  , OTHER, fk_module_WireMessageReply, sensorReading, readingStatus, &fk_module_SensorReading_fields),
    PB_LAST_FIELD
};







/* Check that field information fits in pb_field_t */
#if !defined(PB_FIELD_32BIT)
/* If you get an error here, it means that you need to define PB_FIELD_32BIT
 * compile-time option. You can do that in pb.h or on compiler command line.
 * 
 * The reason you need to do this is that some of your messages contain tag
 * numbers or field sizes that are larger than what can fit in 8 or 16 bit
 * field descriptors.
 */
PB_STATIC_ASSERT((pb_membersize(fk_module_WireMessageQuery, queryCapabilities) < 65536 && pb_membersize(fk_module_WireMessageQuery, querySensorCapabilities) < 65536 && pb_membersize(fk_module_WireMessageQuery, beginTakeReadings) < 65536 && pb_membersize(fk_module_WireMessageQuery, queryReadingStatus) < 65536 && pb_membersize(fk_module_WireMessageReply, error) < 65536 && pb_membersize(fk_module_WireMessageReply, capabilities) < 65536 && pb_membersize(fk_module_WireMessageReply, sensorCapabilities) < 65536 && pb_membersize(fk_module_WireMessageReply, readingStatus) < 65536 && pb_membersize(fk_module_WireMessageReply, sensorReading) < 65536), YOU_MUST_DEFINE_PB_FIELD_32BIT_FOR_MESSAGES_fk_module_QueryCapabilities_fk_module_QuerySensorCapabilities_fk_module_Capabilities_fk_module_SensorCapabilities_fk_module_BeginTakeReadings_fk_module_QueryReadingStatus_fk_module_ReadingStatus_fk_module_SensorReadings_fk_module_SensorReading_fk_module_WireMessageQuery_fk_module_Error_fk_module_WireMessageReply)
#endif

#if !defined(PB_FIELD_16BIT) && !defined(PB_FIELD_32BIT)
/* If you get an error here, it means that you need to define PB_FIELD_16BIT
 * compile-time option. You can do that in pb.h or on compiler command line.
 * 
 * The reason you need to do this is that some of your messages contain tag
 * numbers or field sizes that are larger than what can fit in the default
 * 8 bit descriptors.
 */
PB_STATIC_ASSERT((pb_membersize(fk_module_WireMessageQuery, queryCapabilities) < 256 && pb_membersize(fk_module_WireMessageQuery, querySensorCapabilities) < 256 && pb_membersize(fk_module_WireMessageQuery, beginTakeReadings) < 256 && pb_membersize(fk_module_WireMessageQuery, queryReadingStatus) < 256 && pb_membersize(fk_module_WireMessageReply, error) < 256 && pb_membersize(fk_module_WireMessageReply, capabilities) < 256 && pb_membersize(fk_module_WireMessageReply, sensorCapabilities) < 256 && pb_membersize(fk_module_WireMessageReply, readingStatus) < 256 && pb_membersize(fk_module_WireMessageReply, sensorReading) < 256), YOU_MUST_DEFINE_PB_FIELD_16BIT_FOR_MESSAGES_fk_module_QueryCapabilities_fk_module_QuerySensorCapabilities_fk_module_Capabilities_fk_module_SensorCapabilities_fk_module_BeginTakeReadings_fk_module_QueryReadingStatus_fk_module_ReadingStatus_fk_module_SensorReadings_fk_module_SensorReading_fk_module_WireMessageQuery_fk_module_Error_fk_module_WireMessageReply)
#endif


/* @@protoc_insertion_point(eof) */
