// Code generated by protoc-gen-go. DO NOT EDIT.
// source: master.proto

package masterpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SubDataSetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubDataSetRequest) Reset()         { *m = SubDataSetRequest{} }
func (m *SubDataSetRequest) String() string { return proto.CompactTextString(m) }
func (*SubDataSetRequest) ProtoMessage()    {}
func (*SubDataSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{0}
}

func (m *SubDataSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubDataSetRequest.Unmarshal(m, b)
}
func (m *SubDataSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubDataSetRequest.Marshal(b, m, deterministic)
}
func (m *SubDataSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubDataSetRequest.Merge(m, src)
}
func (m *SubDataSetRequest) XXX_Size() int {
	return xxx_messageInfo_SubDataSetRequest.Size(m)
}
func (m *SubDataSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubDataSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubDataSetRequest proto.InternalMessageInfo

type FileMeta struct {
	IdxInList int64  `protobuf:"varint,1,opt,name=idx_in_list,json=idxInList,proto3" json:"idx_in_list,omitempty"`
	FilePath  string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	// only no
	Records              []*RecordRange `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FileMeta) Reset()         { *m = FileMeta{} }
func (m *FileMeta) String() string { return proto.CompactTextString(m) }
func (*FileMeta) ProtoMessage()    {}
func (*FileMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{1}
}

func (m *FileMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileMeta.Unmarshal(m, b)
}
func (m *FileMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileMeta.Marshal(b, m, deterministic)
}
func (m *FileMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileMeta.Merge(m, src)
}
func (m *FileMeta) XXX_Size() int {
	return xxx_messageInfo_FileMeta.Size(m)
}
func (m *FileMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_FileMeta.DiscardUnknown(m)
}

var xxx_messageInfo_FileMeta proto.InternalMessageInfo

func (m *FileMeta) GetIdxInList() int64 {
	if m != nil {
		return m.IdxInList
	}
	return 0
}

func (m *FileMeta) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *FileMeta) GetRecords() []*RecordRange {
	if m != nil {
		return m.Records
	}
	return nil
}

type SubDataSetResponse struct {
	Ret                  *RPCRet     `protobuf:"bytes,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Metas                []*FileMeta `protobuf:"bytes,2,rep,name=metas,proto3" json:"metas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SubDataSetResponse) Reset()         { *m = SubDataSetResponse{} }
func (m *SubDataSetResponse) String() string { return proto.CompactTextString(m) }
func (*SubDataSetResponse) ProtoMessage()    {}
func (*SubDataSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{2}
}

func (m *SubDataSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubDataSetResponse.Unmarshal(m, b)
}
func (m *SubDataSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubDataSetResponse.Marshal(b, m, deterministic)
}
func (m *SubDataSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubDataSetResponse.Merge(m, src)
}
func (m *SubDataSetResponse) XXX_Size() int {
	return xxx_messageInfo_SubDataSetResponse.Size(m)
}
func (m *SubDataSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubDataSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubDataSetResponse proto.InternalMessageInfo

func (m *SubDataSetResponse) GetRet() *RPCRet {
	if m != nil {
		return m.Ret
	}
	return nil
}

func (m *SubDataSetResponse) GetMetas() []*FileMeta {
	if m != nil {
		return m.Metas
	}
	return nil
}

type DataServer struct {
	PodId                string   `protobuf:"bytes,1,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	Rank                 int32    `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Endpoint             string   `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataServer) Reset()         { *m = DataServer{} }
func (m *DataServer) String() string { return proto.CompactTextString(m) }
func (*DataServer) ProtoMessage()    {}
func (*DataServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{3}
}

func (m *DataServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataServer.Unmarshal(m, b)
}
func (m *DataServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataServer.Marshal(b, m, deterministic)
}
func (m *DataServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataServer.Merge(m, src)
}
func (m *DataServer) XXX_Size() int {
	return xxx_messageInfo_DataServer.Size(m)
}
func (m *DataServer) XXX_DiscardUnknown() {
	xxx_messageInfo_DataServer.DiscardUnknown(m)
}

var xxx_messageInfo_DataServer proto.InternalMessageInfo

func (m *DataServer) GetPodId() string {
	if m != nil {
		return m.PodId
	}
	return ""
}

func (m *DataServer) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *DataServer) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type TrainerRequest struct {
	PodId                string   `protobuf:"bytes,1,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	Rank                 int32    `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Endpoint             string   `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrainerRequest) Reset()         { *m = TrainerRequest{} }
func (m *TrainerRequest) String() string { return proto.CompactTextString(m) }
func (*TrainerRequest) ProtoMessage()    {}
func (*TrainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{4}
}

func (m *TrainerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrainerRequest.Unmarshal(m, b)
}
func (m *TrainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrainerRequest.Marshal(b, m, deterministic)
}
func (m *TrainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrainerRequest.Merge(m, src)
}
func (m *TrainerRequest) XXX_Size() int {
	return xxx_messageInfo_TrainerRequest.Size(m)
}
func (m *TrainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TrainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TrainerRequest proto.InternalMessageInfo

func (m *TrainerRequest) GetPodId() string {
	if m != nil {
		return m.PodId
	}
	return ""
}

func (m *TrainerRequest) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *TrainerRequest) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type Chunk struct {
	DataServer           *DataServer    `protobuf:"bytes,1,opt,name=data_server,json=dataServer,proto3" json:"data_server,omitempty"`
	IdxInList            int64          `protobuf:"varint,2,opt,name=idx_in_list,json=idxInList,proto3" json:"idx_in_list,omitempty"`
	FilePath             string         `protobuf:"bytes,3,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	Records              []*RecordRange `protobuf:"bytes,4,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{5}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetDataServer() *DataServer {
	if m != nil {
		return m.DataServer
	}
	return nil
}

func (m *Chunk) GetIdxInList() int64 {
	if m != nil {
		return m.IdxInList
	}
	return 0
}

func (m *Chunk) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *Chunk) GetRecords() []*RecordRange {
	if m != nil {
		return m.Records
	}
	return nil
}

type ChunksRequest struct {
	Chunks               []*Chunk `protobuf:"bytes,1,rep,name=chunks,proto3" json:"chunks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChunksRequest) Reset()         { *m = ChunksRequest{} }
func (m *ChunksRequest) String() string { return proto.CompactTextString(m) }
func (*ChunksRequest) ProtoMessage()    {}
func (*ChunksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{6}
}

func (m *ChunksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChunksRequest.Unmarshal(m, b)
}
func (m *ChunksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChunksRequest.Marshal(b, m, deterministic)
}
func (m *ChunksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunksRequest.Merge(m, src)
}
func (m *ChunksRequest) XXX_Size() int {
	return xxx_messageInfo_ChunksRequest.Size(m)
}
func (m *ChunksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChunksRequest proto.InternalMessageInfo

func (m *ChunksRequest) GetChunks() []*Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

type TaskMeta struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	EpochNo              int32    `protobuf:"varint,2,opt,name=epoch_no,json=epochNo,proto3" json:"epoch_no,omitempty"`
	Stage                string   `protobuf:"bytes,3,opt,name=Stage,proto3" json:"Stage,omitempty"`
	BatchSize            int32    `protobuf:"varint,4,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskMeta) Reset()         { *m = TaskMeta{} }
func (m *TaskMeta) String() string { return proto.CompactTextString(m) }
func (*TaskMeta) ProtoMessage()    {}
func (*TaskMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{7}
}

func (m *TaskMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskMeta.Unmarshal(m, b)
}
func (m *TaskMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskMeta.Marshal(b, m, deterministic)
}
func (m *TaskMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskMeta.Merge(m, src)
}
func (m *TaskMeta) XXX_Size() int {
	return xxx_messageInfo_TaskMeta.Size(m)
}
func (m *TaskMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskMeta.DiscardUnknown(m)
}

var xxx_messageInfo_TaskMeta proto.InternalMessageInfo

func (m *TaskMeta) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *TaskMeta) GetEpochNo() int32 {
	if m != nil {
		return m.EpochNo
	}
	return 0
}

func (m *TaskMeta) GetStage() string {
	if m != nil {
		return m.Stage
	}
	return ""
}

func (m *TaskMeta) GetBatchSize() int32 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

type Task struct {
	Meta                 *TaskMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Chunks               []*Chunk  `protobuf:"bytes,2,rep,name=chunks,proto3" json:"chunks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{8}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetMeta() *TaskMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Task) GetChunks() []*Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

type Tasks struct {
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tasks) Reset()         { *m = Tasks{} }
func (m *Tasks) String() string { return proto.CompactTextString(m) }
func (*Tasks) ProtoMessage()    {}
func (*Tasks) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{9}
}

func (m *Tasks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tasks.Unmarshal(m, b)
}
func (m *Tasks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tasks.Marshal(b, m, deterministic)
}
func (m *Tasks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tasks.Merge(m, src)
}
func (m *Tasks) XXX_Size() int {
	return xxx_messageInfo_Tasks.Size(m)
}
func (m *Tasks) XXX_DiscardUnknown() {
	xxx_messageInfo_Tasks.DiscardUnknown(m)
}

var xxx_messageInfo_Tasks proto.InternalMessageInfo

func (m *Tasks) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type TaskRequest struct {
	Trainer              *TrainerRequest `protobuf:"bytes,1,opt,name=trainer,proto3" json:"trainer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TaskRequest) Reset()         { *m = TaskRequest{} }
func (m *TaskRequest) String() string { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()    {}
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{10}
}

func (m *TaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRequest.Unmarshal(m, b)
}
func (m *TaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRequest.Marshal(b, m, deterministic)
}
func (m *TaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRequest.Merge(m, src)
}
func (m *TaskRequest) XXX_Size() int {
	return xxx_messageInfo_TaskRequest.Size(m)
}
func (m *TaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRequest proto.InternalMessageInfo

func (m *TaskRequest) GetTrainer() *TrainerRequest {
	if m != nil {
		return m.Trainer
	}
	return nil
}

type TaskResponse struct {
	Ret                  *RPCRet  `protobuf:"bytes,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Tasks                []*Task  `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskResponse) Reset()         { *m = TaskResponse{} }
func (m *TaskResponse) String() string { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()    {}
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{11}
}

func (m *TaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResponse.Unmarshal(m, b)
}
func (m *TaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResponse.Marshal(b, m, deterministic)
}
func (m *TaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResponse.Merge(m, src)
}
func (m *TaskResponse) XXX_Size() int {
	return xxx_messageInfo_TaskResponse.Size(m)
}
func (m *TaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResponse proto.InternalMessageInfo

func (m *TaskResponse) GetRet() *RPCRet {
	if m != nil {
		return m.Ret
	}
	return nil
}

func (m *TaskResponse) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type DataSet struct {
	FileList             string   `protobuf:"bytes,1,opt,name=file_list,json=fileList,proto3" json:"file_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataSet) Reset()         { *m = DataSet{} }
func (m *DataSet) String() string { return proto.CompactTextString(m) }
func (*DataSet) ProtoMessage()    {}
func (*DataSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{12}
}

func (m *DataSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataSet.Unmarshal(m, b)
}
func (m *DataSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataSet.Marshal(b, m, deterministic)
}
func (m *DataSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSet.Merge(m, src)
}
func (m *DataSet) XXX_Size() int {
	return xxx_messageInfo_DataSet.Size(m)
}
func (m *DataSet) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSet.DiscardUnknown(m)
}

var xxx_messageInfo_DataSet proto.InternalMessageInfo

func (m *DataSet) GetFileList() string {
	if m != nil {
		return m.FileList
	}
	return ""
}

type NewEpochRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewEpochRequest) Reset()         { *m = NewEpochRequest{} }
func (m *NewEpochRequest) String() string { return proto.CompactTextString(m) }
func (*NewEpochRequest) ProtoMessage()    {}
func (*NewEpochRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{13}
}

func (m *NewEpochRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewEpochRequest.Unmarshal(m, b)
}
func (m *NewEpochRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewEpochRequest.Marshal(b, m, deterministic)
}
func (m *NewEpochRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewEpochRequest.Merge(m, src)
}
func (m *NewEpochRequest) XXX_Size() int {
	return xxx_messageInfo_NewEpochRequest.Size(m)
}
func (m *NewEpochRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewEpochRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewEpochRequest proto.InternalMessageInfo

type ChunksResponse struct {
	Ret                  *RPCRet  `protobuf:"bytes,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Chunks               []*Chunk `protobuf:"bytes,2,rep,name=chunks,proto3" json:"chunks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChunksResponse) Reset()         { *m = ChunksResponse{} }
func (m *ChunksResponse) String() string { return proto.CompactTextString(m) }
func (*ChunksResponse) ProtoMessage()    {}
func (*ChunksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{14}
}

func (m *ChunksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChunksResponse.Unmarshal(m, b)
}
func (m *ChunksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChunksResponse.Marshal(b, m, deterministic)
}
func (m *ChunksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunksResponse.Merge(m, src)
}
func (m *ChunksResponse) XXX_Size() int {
	return xxx_messageInfo_ChunksResponse.Size(m)
}
func (m *ChunksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChunksResponse proto.InternalMessageInfo

func (m *ChunksResponse) GetRet() *RPCRet {
	if m != nil {
		return m.Ret
	}
	return nil
}

func (m *ChunksResponse) GetChunks() []*Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func init() {
	proto.RegisterType((*SubDataSetRequest)(nil), "master.SubDataSetRequest")
	proto.RegisterType((*FileMeta)(nil), "master.FileMeta")
	proto.RegisterType((*SubDataSetResponse)(nil), "master.SubDataSetResponse")
	proto.RegisterType((*DataServer)(nil), "master.DataServer")
	proto.RegisterType((*TrainerRequest)(nil), "master.TrainerRequest")
	proto.RegisterType((*Chunk)(nil), "master.Chunk")
	proto.RegisterType((*ChunksRequest)(nil), "master.ChunksRequest")
	proto.RegisterType((*TaskMeta)(nil), "master.TaskMeta")
	proto.RegisterType((*Task)(nil), "master.Task")
	proto.RegisterType((*Tasks)(nil), "master.Tasks")
	proto.RegisterType((*TaskRequest)(nil), "master.TaskRequest")
	proto.RegisterType((*TaskResponse)(nil), "master.TaskResponse")
	proto.RegisterType((*DataSet)(nil), "master.DataSet")
	proto.RegisterType((*NewEpochRequest)(nil), "master.NewEpochRequest")
	proto.RegisterType((*ChunksResponse)(nil), "master.ChunksResponse")
}

func init() { proto.RegisterFile("master.proto", fileDescriptor_f9c348dec43a6705) }

var fileDescriptor_f9c348dec43a6705 = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdb, 0x4e, 0x53, 0x4d,
	0x14, 0xa6, 0xe7, 0x76, 0xf5, 0xc0, 0xcf, 0x70, 0x2a, 0xfd, 0xa3, 0x69, 0x76, 0x94, 0x90, 0x18,
	0xab, 0x81, 0x88, 0x17, 0x5e, 0x10, 0x45, 0xc0, 0x26, 0x42, 0xc8, 0x94, 0xc4, 0xe8, 0x85, 0xcd,
	0xb4, 0x33, 0xd2, 0x49, 0xdb, 0x3d, 0xdb, 0x99, 0x29, 0x12, 0x9e, 0xc8, 0xd7, 0xf1, 0x8d, 0xcc,
	0x9e, 0xc3, 0x6e, 0xcb, 0x29, 0xbd, 0xf0, 0xae, 0xeb, 0xfc, 0x7d, 0x6b, 0x7d, 0xb3, 0x0b, 0x95,
	0x31, 0x51, 0x9a, 0xc9, 0x56, 0x24, 0x85, 0x16, 0x28, 0x6f, 0xad, 0x46, 0xa5, 0x2f, 0xc6, 0x63,
	0x11, 0x5a, 0x6f, 0xb0, 0x0a, 0x2b, 0x9d, 0x49, 0xef, 0x23, 0xd1, 0xa4, 0xc3, 0x34, 0x66, 0x3f,
	0x27, 0x4c, 0xe9, 0xe0, 0x0a, 0x8a, 0xc7, 0x7c, 0xc4, 0x4e, 0x99, 0x26, 0xe8, 0x29, 0x94, 0x39,
	0xbd, 0xee, 0xf2, 0xb0, 0x3b, 0xe2, 0x4a, 0xd7, 0x53, 0xcd, 0xd4, 0x4e, 0x06, 0x97, 0x38, 0xbd,
	0x6e, 0x87, 0x9f, 0xb9, 0xd2, 0xe8, 0x7f, 0x28, 0xfd, 0xe0, 0x23, 0xd6, 0x8d, 0x88, 0x1e, 0xd4,
	0xd3, 0xcd, 0xd4, 0x4e, 0x09, 0x17, 0x63, 0xc7, 0x39, 0xd1, 0x03, 0xf4, 0x12, 0x0a, 0x92, 0xf5,
	0x85, 0xa4, 0xaa, 0x9e, 0x69, 0x66, 0x76, 0xca, 0xbb, 0xab, 0x2d, 0x37, 0x1d, 0x1b, 0x37, 0x26,
	0xe1, 0x25, 0xc3, 0x3e, 0x27, 0xf8, 0x0e, 0x68, 0x16, 0x8c, 0x8a, 0x44, 0xa8, 0x18, 0x6a, 0x42,
	0x46, 0x32, 0x3b, 0xb9, 0xbc, 0x5b, 0x4b, 0x1a, 0x9c, 0x1f, 0x62, 0xa6, 0x71, 0x1c, 0x42, 0xdb,
	0x90, 0x1b, 0x33, 0x4d, 0x54, 0x3d, 0x6d, 0x86, 0xfc, 0xd7, 0x72, 0xc4, 0x3d, 0x09, 0x6c, 0xc3,
	0x41, 0x07, 0xc0, 0x36, 0x97, 0x57, 0x4c, 0xa2, 0x75, 0xc8, 0x47, 0x82, 0x76, 0x39, 0x35, 0xad,
	0x4b, 0x38, 0x17, 0x09, 0xda, 0xa6, 0x08, 0x41, 0x56, 0x92, 0x70, 0x68, 0xb8, 0xe4, 0xb0, 0xf9,
	0x8d, 0x1a, 0x50, 0x64, 0x21, 0x8d, 0x04, 0x0f, 0x75, 0x3d, 0x63, 0x39, 0x7a, 0x3b, 0xf8, 0x02,
	0xb5, 0x0b, 0x49, 0x78, 0xc8, 0xa4, 0x5b, 0xdf, 0xbf, 0x6a, 0xfc, 0x3b, 0x05, 0xb9, 0xc3, 0xc1,
	0x24, 0x1c, 0xa2, 0x3d, 0x28, 0x53, 0xa2, 0x49, 0x57, 0x19, 0xe0, 0x6e, 0x13, 0xc8, 0xb3, 0x9c,
	0x52, 0xc2, 0x40, 0xa7, 0xf4, 0x6e, 0x1d, 0x2e, 0xfd, 0xe8, 0xe1, 0x32, 0x0f, 0x1f, 0x2e, 0xbb,
	0xc0, 0xe1, 0xf6, 0xa1, 0x6a, 0x90, 0x2a, 0xbf, 0x82, 0xe7, 0x90, 0xef, 0x1b, 0x47, 0x3d, 0x65,
	0xca, 0xab, 0x1e, 0xac, 0x49, 0xc3, 0x2e, 0x18, 0x28, 0x28, 0x5e, 0x10, 0x35, 0x34, 0x42, 0xdb,
	0x84, 0x82, 0x26, 0x6a, 0x38, 0x5d, 0x5b, 0x3e, 0x36, 0xdb, 0x14, 0x6d, 0x41, 0x91, 0x45, 0xa2,
	0x3f, 0xe8, 0x86, 0xc2, 0xed, 0xae, 0x60, 0xec, 0x33, 0x81, 0xd6, 0x20, 0xd7, 0xd1, 0xe4, 0x92,
	0x39, 0xfc, 0xd6, 0x40, 0x4f, 0x00, 0x7a, 0x44, 0xf7, 0x07, 0x5d, 0xc5, 0x6f, 0x58, 0x3d, 0x6b,
	0x4a, 0x4a, 0xc6, 0xd3, 0xe1, 0x37, 0x2c, 0xe8, 0x40, 0x36, 0x1e, 0x8a, 0x9e, 0x41, 0x36, 0x96,
	0x85, 0x5b, 0x67, 0x22, 0x1a, 0x0f, 0x08, 0x9b, 0xe8, 0x0c, 0x93, 0xf4, 0x63, 0x4c, 0x5e, 0x40,
	0x2e, 0x2e, 0x54, 0x28, 0x80, 0x5c, 0x8c, 0xdb, 0x13, 0xaf, 0xcc, 0xb6, 0xc5, 0x36, 0x14, 0x1c,
	0x40, 0xd9, 0x98, 0x6e, 0x59, 0xaf, 0xa1, 0xa0, 0xad, 0x82, 0x1c, 0x96, 0x8d, 0xa4, 0x68, 0x4e,
	0x58, 0xd8, 0xa7, 0x05, 0x17, 0x50, 0xb1, 0x0d, 0x16, 0x7e, 0x22, 0x09, 0xac, 0xf4, 0xc3, 0xb0,
	0xb6, 0xa1, 0xe0, 0xde, 0x5e, 0x22, 0x8e, 0xe4, 0xcd, 0x3b, 0x71, 0xc4, 0xca, 0x09, 0x56, 0x60,
	0xf9, 0x8c, 0xfd, 0x3a, 0x8a, 0x6f, 0xe0, 0xbf, 0x18, 0x5f, 0xa1, 0xe6, 0x05, 0xb0, 0x30, 0xa4,
	0xc5, 0x36, 0xbb, 0xfb, 0x27, 0x03, 0xf9, 0x53, 0x13, 0x40, 0x07, 0x00, 0x27, 0x4c, 0x1f, 0x8e,
	0x26, 0xc6, 0xda, 0xf0, 0x4d, 0x9d, 0xc3, 0x61, 0x69, 0x6c, 0xde, 0xf1, 0x5b, 0x48, 0xc1, 0x12,
	0xda, 0x83, 0xda, 0x7b, 0x4a, 0xdb, 0x21, 0xd7, 0x9e, 0xe8, 0xf2, 0xfc, 0x2b, 0xd2, 0x8d, 0x5b,
	0x50, 0x83, 0x25, 0xf4, 0x06, 0x8a, 0x9e, 0x2e, 0xda, 0xf4, 0xe9, 0xb7, 0x16, 0x70, 0x4f, 0xd9,
	0x27, 0xa8, 0x9e, 0x30, 0x3d, 0xfd, 0x9e, 0xa1, 0x2d, 0x5f, 0x7b, 0xe7, 0x83, 0xdb, 0x68, 0xdc,
	0x17, 0x4a, 0x50, 0xbf, 0x85, 0x0a, 0x66, 0x91, 0x90, 0xda, 0xae, 0x18, 0xad, 0xcf, 0x2d, 0x4a,
	0x3d, 0x0c, 0x61, 0x1f, 0x0a, 0x27, 0x4c, 0x1b, 0xb1, 0xaf, 0xce, 0x1d, 0xdc, 0x55, 0xac, 0xcd,
	0x3b, 0x93, 0x81, 0xaf, 0xac, 0xbc, 0x8e, 0x79, 0xc8, 0xd5, 0x80, 0x51, 0x54, 0x9d, 0xcd, 0x53,
	0xf7, 0x0c, 0x6a, 0x59, 0x41, 0x1f, 0x49, 0x29, 0xe4, 0x02, 0xf9, 0x1f, 0x2a, 0xdf, 0xa0, 0xf5,
	0xce, 0xe6, 0x44, 0xbd, 0x5e, 0xde, 0xfc, 0x15, 0xed, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x67,
	0x5a, 0xed, 0xe9, 0xb0, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MasterClient interface {
	// Cluster env
	// user_client, computor  -> master
	GetCluster(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterResponse, error)
	// user_client -> Master
	AddInitDataSet(ctx context.Context, in *DataSet, opts ...grpc.CallOption) (*RPCRet, error)
	NewEpoch(ctx context.Context, in *NewEpochRequest, opts ...grpc.CallOption) (*RPCRet, error)
	// data_server - > master
	GetSubDataSet(ctx context.Context, in *SubDataSetRequest, opts ...grpc.CallOption) (*SubDataSetResponse, error)
	ReportChunks(ctx context.Context, in *ChunksRequest, opts ...grpc.CallOption) (*RPCRet, error)
	// executor -> master
	GetTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	TaskFinished(ctx context.Context, in *Tasks, opts ...grpc.CallOption) (*RPCRet, error)
	TaskErrored(ctx context.Context, in *Tasks, opts ...grpc.CallOption) (*RPCRet, error)
}

type masterClient struct {
	cc *grpc.ClientConn
}

func NewMasterClient(cc *grpc.ClientConn) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) GetCluster(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterResponse, error) {
	out := new(ClusterResponse)
	err := c.cc.Invoke(ctx, "/master.Master/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) AddInitDataSet(ctx context.Context, in *DataSet, opts ...grpc.CallOption) (*RPCRet, error) {
	out := new(RPCRet)
	err := c.cc.Invoke(ctx, "/master.Master/AddInitDataSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) NewEpoch(ctx context.Context, in *NewEpochRequest, opts ...grpc.CallOption) (*RPCRet, error) {
	out := new(RPCRet)
	err := c.cc.Invoke(ctx, "/master.Master/NewEpoch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetSubDataSet(ctx context.Context, in *SubDataSetRequest, opts ...grpc.CallOption) (*SubDataSetResponse, error) {
	out := new(SubDataSetResponse)
	err := c.cc.Invoke(ctx, "/master.Master/GetSubDataSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ReportChunks(ctx context.Context, in *ChunksRequest, opts ...grpc.CallOption) (*RPCRet, error) {
	out := new(RPCRet)
	err := c.cc.Invoke(ctx, "/master.Master/ReportChunks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/master.Master/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) TaskFinished(ctx context.Context, in *Tasks, opts ...grpc.CallOption) (*RPCRet, error) {
	out := new(RPCRet)
	err := c.cc.Invoke(ctx, "/master.Master/TaskFinished", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) TaskErrored(ctx context.Context, in *Tasks, opts ...grpc.CallOption) (*RPCRet, error) {
	out := new(RPCRet)
	err := c.cc.Invoke(ctx, "/master.Master/TaskErrored", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServer is the server API for Master service.
type MasterServer interface {
	// Cluster env
	// user_client, computor  -> master
	GetCluster(context.Context, *ClusterRequest) (*ClusterResponse, error)
	// user_client -> Master
	AddInitDataSet(context.Context, *DataSet) (*RPCRet, error)
	NewEpoch(context.Context, *NewEpochRequest) (*RPCRet, error)
	// data_server - > master
	GetSubDataSet(context.Context, *SubDataSetRequest) (*SubDataSetResponse, error)
	ReportChunks(context.Context, *ChunksRequest) (*RPCRet, error)
	// executor -> master
	GetTask(context.Context, *TaskRequest) (*TaskResponse, error)
	TaskFinished(context.Context, *Tasks) (*RPCRet, error)
	TaskErrored(context.Context, *Tasks) (*RPCRet, error)
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetCluster(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_AddInitDataSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).AddInitDataSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/AddInitDataSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).AddInitDataSet(ctx, req.(*DataSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_NewEpoch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewEpochRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).NewEpoch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/NewEpoch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).NewEpoch(ctx, req.(*NewEpochRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetSubDataSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubDataSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetSubDataSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/GetSubDataSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetSubDataSet(ctx, req.(*SubDataSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ReportChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ReportChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/ReportChunks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ReportChunks(ctx, req.(*ChunksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_TaskFinished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tasks)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).TaskFinished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/TaskFinished",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).TaskFinished(ctx, req.(*Tasks))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_TaskErrored_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tasks)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).TaskErrored(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/TaskErrored",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).TaskErrored(ctx, req.(*Tasks))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "master.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCluster",
			Handler:    _Master_GetCluster_Handler,
		},
		{
			MethodName: "AddInitDataSet",
			Handler:    _Master_AddInitDataSet_Handler,
		},
		{
			MethodName: "NewEpoch",
			Handler:    _Master_NewEpoch_Handler,
		},
		{
			MethodName: "GetSubDataSet",
			Handler:    _Master_GetSubDataSet_Handler,
		},
		{
			MethodName: "ReportChunks",
			Handler:    _Master_ReportChunks_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _Master_GetTask_Handler,
		},
		{
			MethodName: "TaskFinished",
			Handler:    _Master_TaskFinished_Handler,
		},
		{
			MethodName: "TaskErrored",
			Handler:    _Master_TaskErrored_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master.proto",
}
