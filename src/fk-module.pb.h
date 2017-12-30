/* Automatically generated nanopb header */
/* Generated by nanopb-0.3.9-dev at Sat Dec 30 11:03:27 2017. */

#ifndef PB_FK_MODULE_FK_MODULE_PB_H_INCLUDED
#define PB_FK_MODULE_FK_MODULE_PB_H_INCLUDED
#include <pb.h>

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Enum definitions */
typedef enum _fk_module_QueryType {
    fk_module_QueryType_QUERY_NONE = 0,
    fk_module_QueryType_QUERY_CAPABILITIES = 1,
    fk_module_QueryType_QUERY_SENSOR_CAPABILITIES = 2,
    fk_module_QueryType_QUERY_BEGIN_TAKE_READINGS = 3,
    fk_module_QueryType_QUERY_READING_STATUS = 4
} fk_module_QueryType;
#define _fk_module_QueryType_MIN fk_module_QueryType_QUERY_NONE
#define _fk_module_QueryType_MAX fk_module_QueryType_QUERY_READING_STATUS
#define _fk_module_QueryType_ARRAYSIZE ((fk_module_QueryType)(fk_module_QueryType_QUERY_READING_STATUS+1))

typedef enum _fk_module_ReplyType {
    fk_module_ReplyType_REPLY_NONE = 0,
    fk_module_ReplyType_REPLY_RETRY = 1,
    fk_module_ReplyType_REPLY_ERROR = 2,
    fk_module_ReplyType_REPLY_CAPABILITIES = 3,
    fk_module_ReplyType_REPLY_SENSOR_CAPABILITIES = 4,
    fk_module_ReplyType_REPLY_READING_STATUS = 5
} fk_module_ReplyType;
#define _fk_module_ReplyType_MIN fk_module_ReplyType_REPLY_NONE
#define _fk_module_ReplyType_MAX fk_module_ReplyType_REPLY_READING_STATUS
#define _fk_module_ReplyType_ARRAYSIZE ((fk_module_ReplyType)(fk_module_ReplyType_REPLY_READING_STATUS+1))

typedef enum _fk_module_ModuleType {
    fk_module_ModuleType_SENSOR = 0,
    fk_module_ModuleType_COMMUNICATIONS = 1
} fk_module_ModuleType;
#define _fk_module_ModuleType_MIN fk_module_ModuleType_SENSOR
#define _fk_module_ModuleType_MAX fk_module_ModuleType_COMMUNICATIONS
#define _fk_module_ModuleType_ARRAYSIZE ((fk_module_ModuleType)(fk_module_ModuleType_COMMUNICATIONS+1))

typedef enum _fk_module_ReadingState {
    fk_module_ReadingState_IDLE = 0,
    fk_module_ReadingState_BEGIN = 1,
    fk_module_ReadingState_BUSY = 2,
    fk_module_ReadingState_DONE = 3
} fk_module_ReadingState;
#define _fk_module_ReadingState_MIN fk_module_ReadingState_IDLE
#define _fk_module_ReadingState_MAX fk_module_ReadingState_DONE
#define _fk_module_ReadingState_ARRAYSIZE ((fk_module_ReadingState)(fk_module_ReadingState_DONE+1))

typedef enum _fk_module_ErrorType {
    fk_module_ErrorType_NONE = 0,
    fk_module_ErrorType_GENERAL = 1,
    fk_module_ErrorType_UNEXPECTED = 2
} fk_module_ErrorType;
#define _fk_module_ErrorType_MIN fk_module_ErrorType_NONE
#define _fk_module_ErrorType_MAX fk_module_ErrorType_UNEXPECTED
#define _fk_module_ErrorType_ARRAYSIZE ((fk_module_ErrorType)(fk_module_ErrorType_UNEXPECTED+1))

/* Struct definitions */
typedef struct _fk_module_QueryReadingStatus {
    char dummy_field;
/* @@protoc_insertion_point(struct:fk_module_QueryReadingStatus) */
} fk_module_QueryReadingStatus;

typedef struct _fk_module_SensorReadings {
    char dummy_field;
/* @@protoc_insertion_point(struct:fk_module_SensorReadings) */
} fk_module_SensorReadings;

typedef struct _fk_module_BeginTakeReadings {
    int32_t index;
    uint32_t callerTime;
/* @@protoc_insertion_point(struct:fk_module_BeginTakeReadings) */
} fk_module_BeginTakeReadings;

typedef struct _fk_module_Capabilities {
    int32_t version;
    fk_module_ModuleType type;
    pb_callback_t name;
    int32_t numberOfSensors;
/* @@protoc_insertion_point(struct:fk_module_Capabilities) */
} fk_module_Capabilities;

typedef struct _fk_module_Error {
    fk_module_ErrorType type;
    pb_callback_t message;
/* @@protoc_insertion_point(struct:fk_module_Error) */
} fk_module_Error;

typedef struct _fk_module_QueryCapabilities {
    int32_t version;
    uint32_t callerTime;
/* @@protoc_insertion_point(struct:fk_module_QueryCapabilities) */
} fk_module_QueryCapabilities;

typedef struct _fk_module_QuerySensorCapabilities {
    int32_t sensor;
/* @@protoc_insertion_point(struct:fk_module_QuerySensorCapabilities) */
} fk_module_QuerySensorCapabilities;

typedef struct _fk_module_ReadingStatus {
    fk_module_ReadingState state;
    uint32_t backoff;
    uint32_t elapsed;
/* @@protoc_insertion_point(struct:fk_module_ReadingStatus) */
} fk_module_ReadingStatus;

typedef struct _fk_module_SensorCapabilities {
    int32_t id;
    pb_callback_t name;
    pb_callback_t unitOfMeasure;
/* @@protoc_insertion_point(struct:fk_module_SensorCapabilities) */
} fk_module_SensorCapabilities;

typedef struct _fk_module_SensorReading {
    int32_t sensor;
    uint32_t time;
    float value;
/* @@protoc_insertion_point(struct:fk_module_SensorReading) */
} fk_module_SensorReading;

typedef struct _fk_module_WireMessageQuery {
    fk_module_QueryType type;
    fk_module_QueryCapabilities queryCapabilities;
    fk_module_QuerySensorCapabilities querySensorCapabilities;
    fk_module_BeginTakeReadings beginTakeReadings;
    fk_module_QueryReadingStatus queryReadingStatus;
/* @@protoc_insertion_point(struct:fk_module_WireMessageQuery) */
} fk_module_WireMessageQuery;

typedef struct _fk_module_WireMessageReply {
    fk_module_ReplyType type;
    fk_module_Error error;
    fk_module_Capabilities capabilities;
    fk_module_SensorCapabilities sensorCapabilities;
    fk_module_ReadingStatus readingStatus;
    fk_module_SensorReading sensorReading;
/* @@protoc_insertion_point(struct:fk_module_WireMessageReply) */
} fk_module_WireMessageReply;

/* Default values for struct fields */

/* Initializer values for message structs */
#define fk_module_QueryCapabilities_init_default {0, 0}
#define fk_module_QuerySensorCapabilities_init_default {0}
#define fk_module_Capabilities_init_default      {0, (fk_module_ModuleType)0, {{NULL}, NULL}, 0}
#define fk_module_SensorCapabilities_init_default {0, {{NULL}, NULL}, {{NULL}, NULL}}
#define fk_module_BeginTakeReadings_init_default {0, 0}
#define fk_module_QueryReadingStatus_init_default {0}
#define fk_module_ReadingStatus_init_default     {(fk_module_ReadingState)0, 0, 0}
#define fk_module_SensorReadings_init_default    {0}
#define fk_module_SensorReading_init_default     {0, 0, 0}
#define fk_module_WireMessageQuery_init_default  {(fk_module_QueryType)0, fk_module_QueryCapabilities_init_default, fk_module_QuerySensorCapabilities_init_default, fk_module_BeginTakeReadings_init_default, fk_module_QueryReadingStatus_init_default}
#define fk_module_Error_init_default             {(fk_module_ErrorType)0, {{NULL}, NULL}}
#define fk_module_WireMessageReply_init_default  {(fk_module_ReplyType)0, fk_module_Error_init_default, fk_module_Capabilities_init_default, fk_module_SensorCapabilities_init_default, fk_module_ReadingStatus_init_default, fk_module_SensorReading_init_default}
#define fk_module_QueryCapabilities_init_zero    {0, 0}
#define fk_module_QuerySensorCapabilities_init_zero {0}
#define fk_module_Capabilities_init_zero         {0, (fk_module_ModuleType)0, {{NULL}, NULL}, 0}
#define fk_module_SensorCapabilities_init_zero   {0, {{NULL}, NULL}, {{NULL}, NULL}}
#define fk_module_BeginTakeReadings_init_zero    {0, 0}
#define fk_module_QueryReadingStatus_init_zero   {0}
#define fk_module_ReadingStatus_init_zero        {(fk_module_ReadingState)0, 0, 0}
#define fk_module_SensorReadings_init_zero       {0}
#define fk_module_SensorReading_init_zero        {0, 0, 0}
#define fk_module_WireMessageQuery_init_zero     {(fk_module_QueryType)0, fk_module_QueryCapabilities_init_zero, fk_module_QuerySensorCapabilities_init_zero, fk_module_BeginTakeReadings_init_zero, fk_module_QueryReadingStatus_init_zero}
#define fk_module_Error_init_zero                {(fk_module_ErrorType)0, {{NULL}, NULL}}
#define fk_module_WireMessageReply_init_zero     {(fk_module_ReplyType)0, fk_module_Error_init_zero, fk_module_Capabilities_init_zero, fk_module_SensorCapabilities_init_zero, fk_module_ReadingStatus_init_zero, fk_module_SensorReading_init_zero}

/* Field tags (for use in manual encoding/decoding) */
#define fk_module_BeginTakeReadings_index_tag    1
#define fk_module_BeginTakeReadings_callerTime_tag 2
#define fk_module_Capabilities_version_tag       1
#define fk_module_Capabilities_type_tag          2
#define fk_module_Capabilities_name_tag          3
#define fk_module_Capabilities_numberOfSensors_tag 4
#define fk_module_Error_type_tag                 1
#define fk_module_Error_message_tag              2
#define fk_module_QueryCapabilities_version_tag  1
#define fk_module_QueryCapabilities_callerTime_tag 2
#define fk_module_QuerySensorCapabilities_sensor_tag 1
#define fk_module_ReadingStatus_state_tag        1
#define fk_module_ReadingStatus_backoff_tag      2
#define fk_module_ReadingStatus_elapsed_tag      3
#define fk_module_SensorCapabilities_id_tag      1
#define fk_module_SensorCapabilities_name_tag    2
#define fk_module_SensorCapabilities_unitOfMeasure_tag 3
#define fk_module_SensorReading_sensor_tag       1
#define fk_module_SensorReading_time_tag         2
#define fk_module_SensorReading_value_tag        3
#define fk_module_WireMessageQuery_type_tag      1
#define fk_module_WireMessageQuery_queryCapabilities_tag 2
#define fk_module_WireMessageQuery_querySensorCapabilities_tag 3
#define fk_module_WireMessageQuery_beginTakeReadings_tag 4
#define fk_module_WireMessageQuery_queryReadingStatus_tag 5
#define fk_module_WireMessageReply_type_tag      1
#define fk_module_WireMessageReply_error_tag     2
#define fk_module_WireMessageReply_capabilities_tag 3
#define fk_module_WireMessageReply_sensorCapabilities_tag 4
#define fk_module_WireMessageReply_readingStatus_tag 5
#define fk_module_WireMessageReply_sensorReading_tag 6

/* Struct field encoding specification for nanopb */
extern const pb_field_t fk_module_QueryCapabilities_fields[3];
extern const pb_field_t fk_module_QuerySensorCapabilities_fields[2];
extern const pb_field_t fk_module_Capabilities_fields[5];
extern const pb_field_t fk_module_SensorCapabilities_fields[4];
extern const pb_field_t fk_module_BeginTakeReadings_fields[3];
extern const pb_field_t fk_module_QueryReadingStatus_fields[1];
extern const pb_field_t fk_module_ReadingStatus_fields[4];
extern const pb_field_t fk_module_SensorReadings_fields[1];
extern const pb_field_t fk_module_SensorReading_fields[4];
extern const pb_field_t fk_module_WireMessageQuery_fields[6];
extern const pb_field_t fk_module_Error_fields[3];
extern const pb_field_t fk_module_WireMessageReply_fields[7];

/* Maximum encoded size of messages (where known) */
#define fk_module_QueryCapabilities_size         17
#define fk_module_QuerySensorCapabilities_size   11
/* fk_module_Capabilities_size depends on runtime parameters */
/* fk_module_SensorCapabilities_size depends on runtime parameters */
#define fk_module_BeginTakeReadings_size         17
#define fk_module_QueryReadingStatus_size        0
#define fk_module_ReadingStatus_size             14
#define fk_module_SensorReadings_size            0
#define fk_module_SensorReading_size             22
#define fk_module_WireMessageQuery_size          55
/* fk_module_Error_size depends on runtime parameters */
#define fk_module_WireMessageReply_size          (60 + fk_module_Error_size + fk_module_Capabilities_size + fk_module_SensorCapabilities_size)

/* Message IDs (where set with "msgid" option) */
#ifdef PB_MSGID

#define FK_MODULE_MESSAGES \


#endif

#ifdef __cplusplus
} /* extern "C" */
#endif
/* @@protoc_insertion_point(eof) */

#endif
