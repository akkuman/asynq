// Copyright 2020 Kentaro Hibino. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: asynq.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TaskMessage is the internal representation of a task with additional
// metadata fields.
// Next ID: 16
type TaskMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type indicates the kind of the task to be performed.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Payload holds data needed to process the task.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Unique identifier for the task.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the queue to which this task belongs.
	Queue string `protobuf:"bytes,4,opt,name=queue,proto3" json:"queue,omitempty"`
	// Max number of retries for this task.
	Retry int32 `protobuf:"varint,5,opt,name=retry,proto3" json:"retry,omitempty"`
	// Number of times this task has been retried so far.
	Retried int32 `protobuf:"varint,6,opt,name=retried,proto3" json:"retried,omitempty"`
	// Error message from the last failure.
	ErrorMsg string `protobuf:"bytes,7,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	// Time of last failure in Unix time,
	// the number of seconds elapsed since January 1, 1970 UTC.
	// Use zero to indicate no last failure.
	LastFailedAt int64 `protobuf:"varint,11,opt,name=last_failed_at,json=lastFailedAt,proto3" json:"last_failed_at,omitempty"`
	// Timeout specifies timeout in seconds.
	// Use zero to indicate no timeout.
	Timeout int64 `protobuf:"varint,8,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Deadline specifies the deadline for the task in Unix time,
	// the number of seconds elapsed since January 1, 1970 UTC.
	// Use zero to indicate no deadline.
	Deadline int64 `protobuf:"varint,9,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// UniqueKey holds the redis key used for uniqueness lock for this task.
	// Empty string indicates that no uniqueness lock was used.
	UniqueKey string `protobuf:"bytes,10,opt,name=unique_key,json=uniqueKey,proto3" json:"unique_key,omitempty"`
	// GroupKey is a name of the group used for task aggregation.
	// This field is optional and empty value means no aggregation for the task.
	GroupKey string `protobuf:"bytes,14,opt,name=group_key,json=groupKey,proto3" json:"group_key,omitempty"`
	// Retention period specified in a number of seconds.
	// The task will be stored in redis as a completed task until the TTL
	// expires.
	Retention int64 `protobuf:"varint,12,opt,name=retention,proto3" json:"retention,omitempty"`
	// Time when the task completed in success in Unix time,
	// the number of seconds elapsed since January 1, 1970 UTC.
	// This field is populated if result_ttl > 0 upon completion.
	CompletedAt int64 `protobuf:"varint,13,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`
	// whether the payload is compressed(zlib)
	Compress *bool `protobuf:"varint,15,opt,name=compress,proto3,oneof" json:"compress,omitempty"`
}

func (x *TaskMessage) Reset() {
	*x = TaskMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskMessage) ProtoMessage() {}

func (x *TaskMessage) ProtoReflect() protoreflect.Message {
	mi := &file_asynq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskMessage.ProtoReflect.Descriptor instead.
func (*TaskMessage) Descriptor() ([]byte, []int) {
	return file_asynq_proto_rawDescGZIP(), []int{0}
}

func (x *TaskMessage) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TaskMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *TaskMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TaskMessage) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *TaskMessage) GetRetry() int32 {
	if x != nil {
		return x.Retry
	}
	return 0
}

func (x *TaskMessage) GetRetried() int32 {
	if x != nil {
		return x.Retried
	}
	return 0
}

func (x *TaskMessage) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *TaskMessage) GetLastFailedAt() int64 {
	if x != nil {
		return x.LastFailedAt
	}
	return 0
}

func (x *TaskMessage) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *TaskMessage) GetDeadline() int64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

func (x *TaskMessage) GetUniqueKey() string {
	if x != nil {
		return x.UniqueKey
	}
	return ""
}

func (x *TaskMessage) GetGroupKey() string {
	if x != nil {
		return x.GroupKey
	}
	return ""
}

func (x *TaskMessage) GetRetention() int64 {
	if x != nil {
		return x.Retention
	}
	return 0
}

func (x *TaskMessage) GetCompletedAt() int64 {
	if x != nil {
		return x.CompletedAt
	}
	return 0
}

func (x *TaskMessage) GetCompress() bool {
	if x != nil && x.Compress != nil {
		return *x.Compress
	}
	return false
}

// ServerInfo holds information about a running server.
type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Host machine the server is running on.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// PID of the server process.
	Pid int32 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	// Unique identifier for this server.
	ServerId string `protobuf:"bytes,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Maximum number of concurrency this server will use.
	Concurrency int32 `protobuf:"varint,4,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	// List of queue names with their priorities.
	// The server will consume tasks from the queues and prioritize
	// queues with higher priority numbers.
	Queues map[string]int32 `protobuf:"bytes,5,rep,name=queues,proto3" json:"queues,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// If set, the server will always consume tasks from a queue with higher
	// priority.
	StrictPriority bool `protobuf:"varint,6,opt,name=strict_priority,json=strictPriority,proto3" json:"strict_priority,omitempty"`
	// Status indicates the status of the server.
	Status string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	// Time this server was started.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Number of workers currently processing tasks.
	ActiveWorkerCount int32 `protobuf:"varint,9,opt,name=active_worker_count,json=activeWorkerCount,proto3" json:"active_worker_count,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_asynq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_asynq_proto_rawDescGZIP(), []int{1}
}

func (x *ServerInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ServerInfo) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ServerInfo) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *ServerInfo) GetConcurrency() int32 {
	if x != nil {
		return x.Concurrency
	}
	return 0
}

func (x *ServerInfo) GetQueues() map[string]int32 {
	if x != nil {
		return x.Queues
	}
	return nil
}

func (x *ServerInfo) GetStrictPriority() bool {
	if x != nil {
		return x.StrictPriority
	}
	return false
}

func (x *ServerInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ServerInfo) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *ServerInfo) GetActiveWorkerCount() int32 {
	if x != nil {
		return x.ActiveWorkerCount
	}
	return 0
}

// WorkerInfo holds information about a running worker.
type WorkerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Host matchine this worker is running on.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// PID of the process in which this worker is running.
	Pid int32 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	// ID of the server in which this worker is running.
	ServerId string `protobuf:"bytes,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// ID of the task this worker is processing.
	TaskId string `protobuf:"bytes,4,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// Type of the task this worker is processing.
	TaskType string `protobuf:"bytes,5,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// Payload of the task this worker is processing.
	TaskPayload []byte `protobuf:"bytes,6,opt,name=task_payload,json=taskPayload,proto3" json:"task_payload,omitempty"`
	// Name of the queue the task the worker is processing belongs.
	Queue string `protobuf:"bytes,7,opt,name=queue,proto3" json:"queue,omitempty"`
	// Time this worker started processing the task.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Deadline by which the worker needs to complete processing
	// the task. If worker exceeds the deadline, the task will fail.
	Deadline *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (x *WorkerInfo) Reset() {
	*x = WorkerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerInfo) ProtoMessage() {}

func (x *WorkerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_asynq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerInfo.ProtoReflect.Descriptor instead.
func (*WorkerInfo) Descriptor() ([]byte, []int) {
	return file_asynq_proto_rawDescGZIP(), []int{2}
}

func (x *WorkerInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *WorkerInfo) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *WorkerInfo) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *WorkerInfo) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *WorkerInfo) GetTaskType() string {
	if x != nil {
		return x.TaskType
	}
	return ""
}

func (x *WorkerInfo) GetTaskPayload() []byte {
	if x != nil {
		return x.TaskPayload
	}
	return nil
}

func (x *WorkerInfo) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *WorkerInfo) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *WorkerInfo) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

// SchedulerEntry holds information about a periodic task registered
// with a scheduler.
type SchedulerEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier of the scheduler entry.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Periodic schedule spec of the entry.
	Spec string `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// Task type of the periodic task.
	TaskType string `protobuf:"bytes,3,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// Task payload of the periodic task.
	TaskPayload []byte `protobuf:"bytes,4,opt,name=task_payload,json=taskPayload,proto3" json:"task_payload,omitempty"`
	// Options used to enqueue the periodic task.
	EnqueueOptions []string `protobuf:"bytes,5,rep,name=enqueue_options,json=enqueueOptions,proto3" json:"enqueue_options,omitempty"`
	// Next time the task will be enqueued.
	NextEnqueueTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=next_enqueue_time,json=nextEnqueueTime,proto3" json:"next_enqueue_time,omitempty"`
	// Last time the task was enqueued.
	// Zero time if task was never enqueued.
	PrevEnqueueTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=prev_enqueue_time,json=prevEnqueueTime,proto3" json:"prev_enqueue_time,omitempty"`
}

func (x *SchedulerEntry) Reset() {
	*x = SchedulerEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynq_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerEntry) ProtoMessage() {}

func (x *SchedulerEntry) ProtoReflect() protoreflect.Message {
	mi := &file_asynq_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerEntry.ProtoReflect.Descriptor instead.
func (*SchedulerEntry) Descriptor() ([]byte, []int) {
	return file_asynq_proto_rawDescGZIP(), []int{3}
}

func (x *SchedulerEntry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SchedulerEntry) GetSpec() string {
	if x != nil {
		return x.Spec
	}
	return ""
}

func (x *SchedulerEntry) GetTaskType() string {
	if x != nil {
		return x.TaskType
	}
	return ""
}

func (x *SchedulerEntry) GetTaskPayload() []byte {
	if x != nil {
		return x.TaskPayload
	}
	return nil
}

func (x *SchedulerEntry) GetEnqueueOptions() []string {
	if x != nil {
		return x.EnqueueOptions
	}
	return nil
}

func (x *SchedulerEntry) GetNextEnqueueTime() *timestamppb.Timestamp {
	if x != nil {
		return x.NextEnqueueTime
	}
	return nil
}

func (x *SchedulerEntry) GetPrevEnqueueTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PrevEnqueueTime
	}
	return nil
}

// SchedulerEnqueueEvent holds information about an enqueue event
// by a scheduler.
type SchedulerEnqueueEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the task that was enqueued.
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// Time the task was enqueued.
	EnqueueTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=enqueue_time,json=enqueueTime,proto3" json:"enqueue_time,omitempty"`
}

func (x *SchedulerEnqueueEvent) Reset() {
	*x = SchedulerEnqueueEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynq_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerEnqueueEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerEnqueueEvent) ProtoMessage() {}

func (x *SchedulerEnqueueEvent) ProtoReflect() protoreflect.Message {
	mi := &file_asynq_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerEnqueueEvent.ProtoReflect.Descriptor instead.
func (*SchedulerEnqueueEvent) Descriptor() ([]byte, []int) {
	return file_asynq_proto_rawDescGZIP(), []int{4}
}

func (x *SchedulerEnqueueEvent) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *SchedulerEnqueueEvent) GetEnqueueTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EnqueueTime
	}
	return nil
}

var File_asynq_proto protoreflect.FileDescriptor

var file_asynq_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x73, 0x79, 0x6e, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x73, 0x79, 0x6e, 0x71, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x03, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x74,
	0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x65, 0x74, 0x72, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x66,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1f, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x22, 0x8f, 0x03,
	0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x35, 0x0a, 0x06, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x73, 0x79, 0x6e, 0x71, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xb1, 0x02, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x70, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x74,
	0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x64,
	0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x22, 0xad, 0x02, 0x0a, 0x0e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x74,
	0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x46, 0x0a, 0x11, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x65, 0x6e, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x70,
	0x72, 0x65, 0x76, 0x5f, 0x65, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x6f, 0x0a, 0x15, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x65, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x65, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x62, 0x69, 0x6b, 0x65, 0x6e, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x71,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_asynq_proto_rawDescOnce sync.Once
	file_asynq_proto_rawDescData = file_asynq_proto_rawDesc
)

func file_asynq_proto_rawDescGZIP() []byte {
	file_asynq_proto_rawDescOnce.Do(func() {
		file_asynq_proto_rawDescData = protoimpl.X.CompressGZIP(file_asynq_proto_rawDescData)
	})
	return file_asynq_proto_rawDescData
}

var file_asynq_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_asynq_proto_goTypes = []interface{}{
	(*TaskMessage)(nil),           // 0: asynq.TaskMessage
	(*ServerInfo)(nil),            // 1: asynq.ServerInfo
	(*WorkerInfo)(nil),            // 2: asynq.WorkerInfo
	(*SchedulerEntry)(nil),        // 3: asynq.SchedulerEntry
	(*SchedulerEnqueueEvent)(nil), // 4: asynq.SchedulerEnqueueEvent
	nil,                           // 5: asynq.ServerInfo.QueuesEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_asynq_proto_depIdxs = []int32{
	5, // 0: asynq.ServerInfo.queues:type_name -> asynq.ServerInfo.QueuesEntry
	6, // 1: asynq.ServerInfo.start_time:type_name -> google.protobuf.Timestamp
	6, // 2: asynq.WorkerInfo.start_time:type_name -> google.protobuf.Timestamp
	6, // 3: asynq.WorkerInfo.deadline:type_name -> google.protobuf.Timestamp
	6, // 4: asynq.SchedulerEntry.next_enqueue_time:type_name -> google.protobuf.Timestamp
	6, // 5: asynq.SchedulerEntry.prev_enqueue_time:type_name -> google.protobuf.Timestamp
	6, // 6: asynq.SchedulerEnqueueEvent.enqueue_time:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_asynq_proto_init() }
func file_asynq_proto_init() {
	if File_asynq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_asynq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_asynq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_asynq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_asynq_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_asynq_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerEnqueueEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_asynq_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_asynq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_asynq_proto_goTypes,
		DependencyIndexes: file_asynq_proto_depIdxs,
		MessageInfos:      file_asynq_proto_msgTypes,
	}.Build()
	File_asynq_proto = out.File
	file_asynq_proto_rawDesc = nil
	file_asynq_proto_goTypes = nil
	file_asynq_proto_depIdxs = nil
}
