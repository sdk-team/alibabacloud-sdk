// This file is auto-generated, don't edit it. Thanks.
package client

import (
  util  "github.com/alibabacloud-go/tea-utils/service"
  rpc  "github.com/alibabacloud-go/tea-rpc/client"
  endpointutil  "github.com/alibabacloud-go/endpoint-util/service"
  "github.com/alibabacloud-go/tea/tea"
)

type QueryRIUtilizationDetailRequest struct {
  RIInstanceId *string `json:"RIInstanceId" xml:"RIInstanceId"`
  InstanceSpec *string `json:"InstanceSpec" xml:"InstanceSpec"`
  RICommodityCode *string `json:"RICommodityCode" xml:"RICommodityCode" require:"true"`
  DeductedInstanceId *string `json:"DeductedInstanceId" xml:"DeductedInstanceId"`
  StartTime *string `json:"StartTime" xml:"StartTime" require:"true"`
  EndTime *string `json:"EndTime" xml:"EndTime" require:"true"`
  PageNum *int `json:"PageNum" xml:"PageNum"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
}

func (s QueryRIUtilizationDetailRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRIUtilizationDetailRequest) GoString() string {
  return s.String()
}

func (s *QueryRIUtilizationDetailRequest) SetRIInstanceId(v string) *QueryRIUtilizationDetailRequest {
  s.RIInstanceId = &v
  return s
}

func (s *QueryRIUtilizationDetailRequest) SetInstanceSpec(v string) *QueryRIUtilizationDetailRequest {
  s.InstanceSpec = &v
  return s
}

func (s *QueryRIUtilizationDetailRequest) SetRICommodityCode(v string) *QueryRIUtilizationDetailRequest {
  s.RICommodityCode = &v
  return s
}

func (s *QueryRIUtilizationDetailRequest) SetDeductedInstanceId(v string) *QueryRIUtilizationDetailRequest {
  s.DeductedInstanceId = &v
  return s
}

func (s *QueryRIUtilizationDetailRequest) SetStartTime(v string) *QueryRIUtilizationDetailRequest {
  s.StartTime = &v
  return s
}

func (s *QueryRIUtilizationDetailRequest) SetEndTime(v string) *QueryRIUtilizationDetailRequest {
  s.EndTime = &v
  return s
}

func (s *QueryRIUtilizationDetailRequest) SetPageNum(v int) *QueryRIUtilizationDetailRequest {
  s.PageNum = &v
  return s
}

func (s *QueryRIUtilizationDetailRequest) SetPageSize(v int) *QueryRIUtilizationDetailRequest {
  s.PageSize = &v
  return s
}

type QueryRIUtilizationDetailResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryRIUtilizationDetailResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryRIUtilizationDetailResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRIUtilizationDetailResponse) GoString() string {
  return s.String()
}

func (s *QueryRIUtilizationDetailResponse) SetRequestId(v string) *QueryRIUtilizationDetailResponse {
  s.RequestId = &v
  return s
}

func (s *QueryRIUtilizationDetailResponse) SetSuccess(v bool) *QueryRIUtilizationDetailResponse {
  s.Success = &v
  return s
}

func (s *QueryRIUtilizationDetailResponse) SetCode(v string) *QueryRIUtilizationDetailResponse {
  s.Code = &v
  return s
}

func (s *QueryRIUtilizationDetailResponse) SetMessage(v string) *QueryRIUtilizationDetailResponse {
  s.Message = &v
  return s
}

func (s *QueryRIUtilizationDetailResponse) SetData(v *QueryRIUtilizationDetailResponseData) *QueryRIUtilizationDetailResponse {
  s.Data = v
  return s
}

type QueryRIUtilizationDetailResponseData struct {
  PageNum *int64 `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int64 `json:"PageSize" xml:"PageSize" require:"true"`
  TotalCount *int64 `json:"TotalCount" xml:"TotalCount" require:"true"`
  DetailList *QueryRIUtilizationDetailResponseDataDetailList `json:"DetailList" xml:"DetailList" require:"true" type:"Struct"`
}

func (s QueryRIUtilizationDetailResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryRIUtilizationDetailResponseData) GoString() string {
  return s.String()
}

func (s *QueryRIUtilizationDetailResponseData) SetPageNum(v int64) *QueryRIUtilizationDetailResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryRIUtilizationDetailResponseData) SetPageSize(v int64) *QueryRIUtilizationDetailResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryRIUtilizationDetailResponseData) SetTotalCount(v int64) *QueryRIUtilizationDetailResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryRIUtilizationDetailResponseData) SetDetailList(v *QueryRIUtilizationDetailResponseDataDetailList) *QueryRIUtilizationDetailResponseData {
  s.DetailList = v
  return s
}

type QueryRIUtilizationDetailResponseDataDetailList struct {
  DetailList []*QueryRIUtilizationDetailResponseDataDetailListDetailList `json:"DetailList" xml:"DetailList" require:"true" type:"Repeated"`
}

func (s QueryRIUtilizationDetailResponseDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s QueryRIUtilizationDetailResponseDataDetailList) GoString() string {
  return s.String()
}

func (s *QueryRIUtilizationDetailResponseDataDetailList) SetDetailList(v []*QueryRIUtilizationDetailResponseDataDetailListDetailList) *QueryRIUtilizationDetailResponseDataDetailList {
  s.DetailList = v
  return s
}

type QueryRIUtilizationDetailResponseDataDetailListDetailList struct     {
  RIInstanceId *string `json:"RIInstanceId" xml:"RIInstanceId" require:"true"`
  InstanceSpec *string `json:"InstanceSpec" xml:"InstanceSpec" require:"true"`
  DeductedInstanceId *string `json:"DeductedInstanceId" xml:"DeductedInstanceId" require:"true"`
  DeductedCommodityCode *string `json:"DeductedCommodityCode" xml:"DeductedCommodityCode" require:"true"`
  DeductDate *string `json:"DeductDate" xml:"DeductDate" require:"true"`
  DeductHours *string `json:"DeductHours" xml:"DeductHours" require:"true"`
  DeductedProductDetail *string `json:"DeductedProductDetail" xml:"DeductedProductDetail" require:"true"`
  DeductQuantity *float32 `json:"DeductQuantity" xml:"DeductQuantity" require:"true"`
  DeductFactorTotal *float32 `json:"DeductFactorTotal" xml:"DeductFactorTotal" require:"true"`
}

func (s QueryRIUtilizationDetailResponseDataDetailListDetailList) String() string {
  return tea.Prettify(s)
}

func (s QueryRIUtilizationDetailResponseDataDetailListDetailList) GoString() string {
  return s.String()
}

func (s *QueryRIUtilizationDetailResponseDataDetailListDetailList) SetRIInstanceId(v string) *QueryRIUtilizationDetailResponseDataDetailListDetailList {
  s.RIInstanceId = &v
  return s
}

func (s *QueryRIUtilizationDetailResponseDataDetailListDetailList) SetInstanceSpec(v string) *QueryRIUtilizationDetailResponseDataDetailListDetailList {
  s.InstanceSpec = &v
  return s
}

func (s *QueryRIUtilizationDetailResponseDataDetailListDetailList) SetDeductedInstanceId(v string) *QueryRIUtilizationDetailResponseDataDetailListDetailList {
  s.DeductedInstanceId = &v
  return s
}

func (s *QueryRIUtilizationDetailResponseDataDetailListDetailList) SetDeductedCommodityCode(v string) *QueryRIUtilizationDetailResponseDataDetailListDetailList {
  s.DeductedCommodityCode = &v
  return s
}

func (s *QueryRIUtilizationDetailResponseDataDetailListDetailList) SetDeductDate(v string) *QueryRIUtilizationDetailResponseDataDetailListDetailList {
  s.DeductDate = &v
  return s
}

func (s *QueryRIUtilizationDetailResponseDataDetailListDetailList) SetDeductHours(v string) *QueryRIUtilizationDetailResponseDataDetailListDetailList {
  s.DeductHours = &v
  return s
}

func (s *QueryRIUtilizationDetailResponseDataDetailListDetailList) SetDeductedProductDetail(v string) *QueryRIUtilizationDetailResponseDataDetailListDetailList {
  s.DeductedProductDetail = &v
  return s
}

func (s *QueryRIUtilizationDetailResponseDataDetailListDetailList) SetDeductQuantity(v float32) *QueryRIUtilizationDetailResponseDataDetailListDetailList {
  s.DeductQuantity = &v
  return s
}

func (s *QueryRIUtilizationDetailResponseDataDetailListDetailList) SetDeductFactorTotal(v float32) *QueryRIUtilizationDetailResponseDataDetailListDetailList {
  s.DeductFactorTotal = &v
  return s
}

type QueryBillToOSSSubscriptionRequest struct {
}

func (s QueryBillToOSSSubscriptionRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBillToOSSSubscriptionRequest) GoString() string {
  return s.String()
}

type QueryBillToOSSSubscriptionResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryBillToOSSSubscriptionResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryBillToOSSSubscriptionResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBillToOSSSubscriptionResponse) GoString() string {
  return s.String()
}

func (s *QueryBillToOSSSubscriptionResponse) SetRequestId(v string) *QueryBillToOSSSubscriptionResponse {
  s.RequestId = &v
  return s
}

func (s *QueryBillToOSSSubscriptionResponse) SetSuccess(v bool) *QueryBillToOSSSubscriptionResponse {
  s.Success = &v
  return s
}

func (s *QueryBillToOSSSubscriptionResponse) SetCode(v string) *QueryBillToOSSSubscriptionResponse {
  s.Code = &v
  return s
}

func (s *QueryBillToOSSSubscriptionResponse) SetMessage(v string) *QueryBillToOSSSubscriptionResponse {
  s.Message = &v
  return s
}

func (s *QueryBillToOSSSubscriptionResponse) SetData(v *QueryBillToOSSSubscriptionResponseData) *QueryBillToOSSSubscriptionResponse {
  s.Data = v
  return s
}

type QueryBillToOSSSubscriptionResponseData struct {
  AccountID *string `json:"AccountID" xml:"AccountID" require:"true"`
  AccountName *string `json:"AccountName" xml:"AccountName" require:"true"`
  Items *QueryBillToOSSSubscriptionResponseDataItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s QueryBillToOSSSubscriptionResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryBillToOSSSubscriptionResponseData) GoString() string {
  return s.String()
}

func (s *QueryBillToOSSSubscriptionResponseData) SetAccountID(v string) *QueryBillToOSSSubscriptionResponseData {
  s.AccountID = &v
  return s
}

func (s *QueryBillToOSSSubscriptionResponseData) SetAccountName(v string) *QueryBillToOSSSubscriptionResponseData {
  s.AccountName = &v
  return s
}

func (s *QueryBillToOSSSubscriptionResponseData) SetItems(v *QueryBillToOSSSubscriptionResponseDataItems) *QueryBillToOSSSubscriptionResponseData {
  s.Items = v
  return s
}

type QueryBillToOSSSubscriptionResponseDataItems struct {
  Item []*QueryBillToOSSSubscriptionResponseDataItemsItem `json:"Item" xml:"Item" require:"true" type:"Repeated"`
}

func (s QueryBillToOSSSubscriptionResponseDataItems) String() string {
  return tea.Prettify(s)
}

func (s QueryBillToOSSSubscriptionResponseDataItems) GoString() string {
  return s.String()
}

func (s *QueryBillToOSSSubscriptionResponseDataItems) SetItem(v []*QueryBillToOSSSubscriptionResponseDataItemsItem) *QueryBillToOSSSubscriptionResponseDataItems {
  s.Item = v
  return s
}

type QueryBillToOSSSubscriptionResponseDataItemsItem struct     {
  SubscribeType *string `json:"SubscribeType" xml:"SubscribeType" require:"true"`
  SubscribeBucket *string `json:"SubscribeBucket" xml:"SubscribeBucket" require:"true"`
  BucketOwnerId *int64 `json:"BucketOwnerId" xml:"BucketOwnerId" require:"true"`
  SubscribeTime *string `json:"SubscribeTime" xml:"SubscribeTime" require:"true"`
  SubscribeLanguage *string `json:"SubscribeLanguage" xml:"SubscribeLanguage" require:"true"`
  MultAccountRelSubscribe *string `json:"MultAccountRelSubscribe" xml:"MultAccountRelSubscribe" require:"true"`
}

func (s QueryBillToOSSSubscriptionResponseDataItemsItem) String() string {
  return tea.Prettify(s)
}

func (s QueryBillToOSSSubscriptionResponseDataItemsItem) GoString() string {
  return s.String()
}

func (s *QueryBillToOSSSubscriptionResponseDataItemsItem) SetSubscribeType(v string) *QueryBillToOSSSubscriptionResponseDataItemsItem {
  s.SubscribeType = &v
  return s
}

func (s *QueryBillToOSSSubscriptionResponseDataItemsItem) SetSubscribeBucket(v string) *QueryBillToOSSSubscriptionResponseDataItemsItem {
  s.SubscribeBucket = &v
  return s
}

func (s *QueryBillToOSSSubscriptionResponseDataItemsItem) SetBucketOwnerId(v int64) *QueryBillToOSSSubscriptionResponseDataItemsItem {
  s.BucketOwnerId = &v
  return s
}

func (s *QueryBillToOSSSubscriptionResponseDataItemsItem) SetSubscribeTime(v string) *QueryBillToOSSSubscriptionResponseDataItemsItem {
  s.SubscribeTime = &v
  return s
}

func (s *QueryBillToOSSSubscriptionResponseDataItemsItem) SetSubscribeLanguage(v string) *QueryBillToOSSSubscriptionResponseDataItemsItem {
  s.SubscribeLanguage = &v
  return s
}

func (s *QueryBillToOSSSubscriptionResponseDataItemsItem) SetMultAccountRelSubscribe(v string) *QueryBillToOSSSubscriptionResponseDataItemsItem {
  s.MultAccountRelSubscribe = &v
  return s
}

type QueryAccountBillRequest struct {
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  PageNum *int `json:"PageNum" xml:"PageNum"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
  OwnerID *int64 `json:"OwnerID" xml:"OwnerID"`
  IsGroupByProduct *bool `json:"IsGroupByProduct" xml:"IsGroupByProduct"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode"`
}

func (s QueryAccountBillRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountBillRequest) GoString() string {
  return s.String()
}

func (s *QueryAccountBillRequest) SetBillingCycle(v string) *QueryAccountBillRequest {
  s.BillingCycle = &v
  return s
}

func (s *QueryAccountBillRequest) SetPageNum(v int) *QueryAccountBillRequest {
  s.PageNum = &v
  return s
}

func (s *QueryAccountBillRequest) SetPageSize(v int) *QueryAccountBillRequest {
  s.PageSize = &v
  return s
}

func (s *QueryAccountBillRequest) SetOwnerID(v int64) *QueryAccountBillRequest {
  s.OwnerID = &v
  return s
}

func (s *QueryAccountBillRequest) SetIsGroupByProduct(v bool) *QueryAccountBillRequest {
  s.IsGroupByProduct = &v
  return s
}

func (s *QueryAccountBillRequest) SetProductCode(v string) *QueryAccountBillRequest {
  s.ProductCode = &v
  return s
}

type QueryAccountBillResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryAccountBillResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryAccountBillResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountBillResponse) GoString() string {
  return s.String()
}

func (s *QueryAccountBillResponse) SetRequestId(v string) *QueryAccountBillResponse {
  s.RequestId = &v
  return s
}

func (s *QueryAccountBillResponse) SetSuccess(v bool) *QueryAccountBillResponse {
  s.Success = &v
  return s
}

func (s *QueryAccountBillResponse) SetCode(v string) *QueryAccountBillResponse {
  s.Code = &v
  return s
}

func (s *QueryAccountBillResponse) SetMessage(v string) *QueryAccountBillResponse {
  s.Message = &v
  return s
}

func (s *QueryAccountBillResponse) SetData(v *QueryAccountBillResponseData) *QueryAccountBillResponse {
  s.Data = v
  return s
}

type QueryAccountBillResponseData struct {
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  AccountID *string `json:"AccountID" xml:"AccountID" require:"true"`
  AccountName *string `json:"AccountName" xml:"AccountName" require:"true"`
  TotalCount *int `json:"TotalCount" xml:"TotalCount" require:"true"`
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  Items *QueryAccountBillResponseDataItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s QueryAccountBillResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountBillResponseData) GoString() string {
  return s.String()
}

func (s *QueryAccountBillResponseData) SetBillingCycle(v string) *QueryAccountBillResponseData {
  s.BillingCycle = &v
  return s
}

func (s *QueryAccountBillResponseData) SetAccountID(v string) *QueryAccountBillResponseData {
  s.AccountID = &v
  return s
}

func (s *QueryAccountBillResponseData) SetAccountName(v string) *QueryAccountBillResponseData {
  s.AccountName = &v
  return s
}

func (s *QueryAccountBillResponseData) SetTotalCount(v int) *QueryAccountBillResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryAccountBillResponseData) SetPageNum(v int) *QueryAccountBillResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryAccountBillResponseData) SetPageSize(v int) *QueryAccountBillResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryAccountBillResponseData) SetItems(v *QueryAccountBillResponseDataItems) *QueryAccountBillResponseData {
  s.Items = v
  return s
}

type QueryAccountBillResponseDataItems struct {
  Item []*QueryAccountBillResponseDataItemsItem `json:"Item" xml:"Item" require:"true" type:"Repeated"`
}

func (s QueryAccountBillResponseDataItems) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountBillResponseDataItems) GoString() string {
  return s.String()
}

func (s *QueryAccountBillResponseDataItems) SetItem(v []*QueryAccountBillResponseDataItemsItem) *QueryAccountBillResponseDataItems {
  s.Item = v
  return s
}

type QueryAccountBillResponseDataItemsItem struct     {
  CostUnit *string `json:"CostUnit" xml:"CostUnit" require:"true"`
  OwnerID *string `json:"OwnerID" xml:"OwnerID" require:"true"`
  PretaxGrossAmount *float32 `json:"PretaxGrossAmount" xml:"PretaxGrossAmount" require:"true"`
  InvoiceDiscount *float32 `json:"InvoiceDiscount" xml:"InvoiceDiscount" require:"true"`
  DeductedByCoupons *float32 `json:"DeductedByCoupons" xml:"DeductedByCoupons" require:"true"`
  PretaxAmount *float32 `json:"PretaxAmount" xml:"PretaxAmount" require:"true"`
  DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons" xml:"DeductedByCashCoupons" require:"true"`
  DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard" xml:"DeductedByPrepaidCard" require:"true"`
  PaymentAmount *float32 `json:"PaymentAmount" xml:"PaymentAmount" require:"true"`
  OutstandingAmount *float32 `json:"OutstandingAmount" xml:"OutstandingAmount" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  OwnerName *string `json:"OwnerName" xml:"OwnerName" require:"true"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductName *string `json:"ProductName" xml:"ProductName" require:"true"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
}

func (s QueryAccountBillResponseDataItemsItem) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountBillResponseDataItemsItem) GoString() string {
  return s.String()
}

func (s *QueryAccountBillResponseDataItemsItem) SetCostUnit(v string) *QueryAccountBillResponseDataItemsItem {
  s.CostUnit = &v
  return s
}

func (s *QueryAccountBillResponseDataItemsItem) SetOwnerID(v string) *QueryAccountBillResponseDataItemsItem {
  s.OwnerID = &v
  return s
}

func (s *QueryAccountBillResponseDataItemsItem) SetPretaxGrossAmount(v float32) *QueryAccountBillResponseDataItemsItem {
  s.PretaxGrossAmount = &v
  return s
}

func (s *QueryAccountBillResponseDataItemsItem) SetInvoiceDiscount(v float32) *QueryAccountBillResponseDataItemsItem {
  s.InvoiceDiscount = &v
  return s
}

func (s *QueryAccountBillResponseDataItemsItem) SetDeductedByCoupons(v float32) *QueryAccountBillResponseDataItemsItem {
  s.DeductedByCoupons = &v
  return s
}

func (s *QueryAccountBillResponseDataItemsItem) SetPretaxAmount(v float32) *QueryAccountBillResponseDataItemsItem {
  s.PretaxAmount = &v
  return s
}

func (s *QueryAccountBillResponseDataItemsItem) SetDeductedByCashCoupons(v float32) *QueryAccountBillResponseDataItemsItem {
  s.DeductedByCashCoupons = &v
  return s
}

func (s *QueryAccountBillResponseDataItemsItem) SetDeductedByPrepaidCard(v float32) *QueryAccountBillResponseDataItemsItem {
  s.DeductedByPrepaidCard = &v
  return s
}

func (s *QueryAccountBillResponseDataItemsItem) SetPaymentAmount(v float32) *QueryAccountBillResponseDataItemsItem {
  s.PaymentAmount = &v
  return s
}

func (s *QueryAccountBillResponseDataItemsItem) SetOutstandingAmount(v float32) *QueryAccountBillResponseDataItemsItem {
  s.OutstandingAmount = &v
  return s
}

func (s *QueryAccountBillResponseDataItemsItem) SetCurrency(v string) *QueryAccountBillResponseDataItemsItem {
  s.Currency = &v
  return s
}

func (s *QueryAccountBillResponseDataItemsItem) SetOwnerName(v string) *QueryAccountBillResponseDataItemsItem {
  s.OwnerName = &v
  return s
}

func (s *QueryAccountBillResponseDataItemsItem) SetProductCode(v string) *QueryAccountBillResponseDataItemsItem {
  s.ProductCode = &v
  return s
}

func (s *QueryAccountBillResponseDataItemsItem) SetProductName(v string) *QueryAccountBillResponseDataItemsItem {
  s.ProductName = &v
  return s
}

func (s *QueryAccountBillResponseDataItemsItem) SetSubscriptionType(v string) *QueryAccountBillResponseDataItemsItem {
  s.SubscriptionType = &v
  return s
}

type CreateCostUnitRequest struct {
  UnitEntityList []*CreateCostUnitRequestUnitEntityList `json:"UnitEntityList" xml:"UnitEntityList" type:"Repeated"`
}

func (s CreateCostUnitRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateCostUnitRequest) GoString() string {
  return s.String()
}

func (s *CreateCostUnitRequest) SetUnitEntityList(v []*CreateCostUnitRequestUnitEntityList) *CreateCostUnitRequest {
  s.UnitEntityList = v
  return s
}

type CreateCostUnitRequestUnitEntityList struct     {
  OwnerUid *int64 `json:"OwnerUid" xml:"OwnerUid" require:"true"`
  ParentUnitId *int64 `json:"ParentUnitId" xml:"ParentUnitId" require:"true"`
  UnitName *string `json:"UnitName" xml:"UnitName" require:"true"`
}

func (s CreateCostUnitRequestUnitEntityList) String() string {
  return tea.Prettify(s)
}

func (s CreateCostUnitRequestUnitEntityList) GoString() string {
  return s.String()
}

func (s *CreateCostUnitRequestUnitEntityList) SetOwnerUid(v int64) *CreateCostUnitRequestUnitEntityList {
  s.OwnerUid = &v
  return s
}

func (s *CreateCostUnitRequestUnitEntityList) SetParentUnitId(v int64) *CreateCostUnitRequestUnitEntityList {
  s.ParentUnitId = &v
  return s
}

func (s *CreateCostUnitRequestUnitEntityList) SetUnitName(v string) *CreateCostUnitRequestUnitEntityList {
  s.UnitName = &v
  return s
}

type CreateCostUnitResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *CreateCostUnitResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s CreateCostUnitResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateCostUnitResponse) GoString() string {
  return s.String()
}

func (s *CreateCostUnitResponse) SetRequestId(v string) *CreateCostUnitResponse {
  s.RequestId = &v
  return s
}

func (s *CreateCostUnitResponse) SetSuccess(v bool) *CreateCostUnitResponse {
  s.Success = &v
  return s
}

func (s *CreateCostUnitResponse) SetCode(v string) *CreateCostUnitResponse {
  s.Code = &v
  return s
}

func (s *CreateCostUnitResponse) SetMessage(v string) *CreateCostUnitResponse {
  s.Message = &v
  return s
}

func (s *CreateCostUnitResponse) SetData(v *CreateCostUnitResponseData) *CreateCostUnitResponse {
  s.Data = v
  return s
}

type CreateCostUnitResponseData struct {
  CostUnitDtoList []*CreateCostUnitResponseDataCostUnitDtoList `json:"CostUnitDtoList" xml:"CostUnitDtoList" require:"true" type:"Repeated"`
}

func (s CreateCostUnitResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateCostUnitResponseData) GoString() string {
  return s.String()
}

func (s *CreateCostUnitResponseData) SetCostUnitDtoList(v []*CreateCostUnitResponseDataCostUnitDtoList) *CreateCostUnitResponseData {
  s.CostUnitDtoList = v
  return s
}

type CreateCostUnitResponseDataCostUnitDtoList struct     {
  OwnerUid *int64 `json:"OwnerUid" xml:"OwnerUid" require:"true"`
  ParentUnitId *int64 `json:"ParentUnitId" xml:"ParentUnitId" require:"true"`
  UnitId *int64 `json:"UnitId" xml:"UnitId" require:"true"`
  UnitName *string `json:"UnitName" xml:"UnitName" require:"true"`
}

func (s CreateCostUnitResponseDataCostUnitDtoList) String() string {
  return tea.Prettify(s)
}

func (s CreateCostUnitResponseDataCostUnitDtoList) GoString() string {
  return s.String()
}

func (s *CreateCostUnitResponseDataCostUnitDtoList) SetOwnerUid(v int64) *CreateCostUnitResponseDataCostUnitDtoList {
  s.OwnerUid = &v
  return s
}

func (s *CreateCostUnitResponseDataCostUnitDtoList) SetParentUnitId(v int64) *CreateCostUnitResponseDataCostUnitDtoList {
  s.ParentUnitId = &v
  return s
}

func (s *CreateCostUnitResponseDataCostUnitDtoList) SetUnitId(v int64) *CreateCostUnitResponseDataCostUnitDtoList {
  s.UnitId = &v
  return s
}

func (s *CreateCostUnitResponseDataCostUnitDtoList) SetUnitName(v string) *CreateCostUnitResponseDataCostUnitDtoList {
  s.UnitName = &v
  return s
}

type ModifyCostUnitRequest struct {
  UnitEntityList []*ModifyCostUnitRequestUnitEntityList `json:"UnitEntityList" xml:"UnitEntityList" type:"Repeated"`
}

func (s ModifyCostUnitRequest) String() string {
  return tea.Prettify(s)
}

func (s ModifyCostUnitRequest) GoString() string {
  return s.String()
}

func (s *ModifyCostUnitRequest) SetUnitEntityList(v []*ModifyCostUnitRequestUnitEntityList) *ModifyCostUnitRequest {
  s.UnitEntityList = v
  return s
}

type ModifyCostUnitRequestUnitEntityList struct     {
  OwnerUid *int64 `json:"OwnerUid" xml:"OwnerUid" require:"true"`
  UnitId *int64 `json:"UnitId" xml:"UnitId" require:"true"`
  NewUnitName *string `json:"NewUnitName" xml:"NewUnitName" require:"true"`
}

func (s ModifyCostUnitRequestUnitEntityList) String() string {
  return tea.Prettify(s)
}

func (s ModifyCostUnitRequestUnitEntityList) GoString() string {
  return s.String()
}

func (s *ModifyCostUnitRequestUnitEntityList) SetOwnerUid(v int64) *ModifyCostUnitRequestUnitEntityList {
  s.OwnerUid = &v
  return s
}

func (s *ModifyCostUnitRequestUnitEntityList) SetUnitId(v int64) *ModifyCostUnitRequestUnitEntityList {
  s.UnitId = &v
  return s
}

func (s *ModifyCostUnitRequestUnitEntityList) SetNewUnitName(v string) *ModifyCostUnitRequestUnitEntityList {
  s.NewUnitName = &v
  return s
}

type ModifyCostUnitResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data []*ModifyCostUnitResponseData `json:"Data" xml:"Data" require:"true" type:"Repeated"`
}

func (s ModifyCostUnitResponse) String() string {
  return tea.Prettify(s)
}

func (s ModifyCostUnitResponse) GoString() string {
  return s.String()
}

func (s *ModifyCostUnitResponse) SetRequestId(v string) *ModifyCostUnitResponse {
  s.RequestId = &v
  return s
}

func (s *ModifyCostUnitResponse) SetSuccess(v bool) *ModifyCostUnitResponse {
  s.Success = &v
  return s
}

func (s *ModifyCostUnitResponse) SetCode(v string) *ModifyCostUnitResponse {
  s.Code = &v
  return s
}

func (s *ModifyCostUnitResponse) SetMessage(v string) *ModifyCostUnitResponse {
  s.Message = &v
  return s
}

func (s *ModifyCostUnitResponse) SetData(v []*ModifyCostUnitResponseData) *ModifyCostUnitResponse {
  s.Data = v
  return s
}

type ModifyCostUnitResponseData struct     {
  OwnerUid *int64 `json:"OwnerUid" xml:"OwnerUid" require:"true"`
  UnitId *int64 `json:"UnitId" xml:"UnitId" require:"true"`
  IsSuccess *bool `json:"IsSuccess" xml:"IsSuccess" require:"true"`
}

func (s ModifyCostUnitResponseData) String() string {
  return tea.Prettify(s)
}

func (s ModifyCostUnitResponseData) GoString() string {
  return s.String()
}

func (s *ModifyCostUnitResponseData) SetOwnerUid(v int64) *ModifyCostUnitResponseData {
  s.OwnerUid = &v
  return s
}

func (s *ModifyCostUnitResponseData) SetUnitId(v int64) *ModifyCostUnitResponseData {
  s.UnitId = &v
  return s
}

func (s *ModifyCostUnitResponseData) SetIsSuccess(v bool) *ModifyCostUnitResponseData {
  s.IsSuccess = &v
  return s
}

type QueryCostUnitRequest struct {
  OwnerUid *int64 `json:"OwnerUid" xml:"OwnerUid" require:"true"`
  ParentUnitId *int64 `json:"ParentUnitId" xml:"ParentUnitId" require:"true"`
  PageNum *int `json:"PageNum" xml:"PageNum"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
}

func (s QueryCostUnitRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCostUnitRequest) GoString() string {
  return s.String()
}

func (s *QueryCostUnitRequest) SetOwnerUid(v int64) *QueryCostUnitRequest {
  s.OwnerUid = &v
  return s
}

func (s *QueryCostUnitRequest) SetParentUnitId(v int64) *QueryCostUnitRequest {
  s.ParentUnitId = &v
  return s
}

func (s *QueryCostUnitRequest) SetPageNum(v int) *QueryCostUnitRequest {
  s.PageNum = &v
  return s
}

func (s *QueryCostUnitRequest) SetPageSize(v int) *QueryCostUnitRequest {
  s.PageSize = &v
  return s
}

type QueryCostUnitResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryCostUnitResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryCostUnitResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCostUnitResponse) GoString() string {
  return s.String()
}

func (s *QueryCostUnitResponse) SetRequestId(v string) *QueryCostUnitResponse {
  s.RequestId = &v
  return s
}

func (s *QueryCostUnitResponse) SetSuccess(v bool) *QueryCostUnitResponse {
  s.Success = &v
  return s
}

func (s *QueryCostUnitResponse) SetCode(v string) *QueryCostUnitResponse {
  s.Code = &v
  return s
}

func (s *QueryCostUnitResponse) SetMessage(v string) *QueryCostUnitResponse {
  s.Message = &v
  return s
}

func (s *QueryCostUnitResponse) SetData(v *QueryCostUnitResponseData) *QueryCostUnitResponse {
  s.Data = v
  return s
}

type QueryCostUnitResponseData struct {
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  TotalCount *int `json:"TotalCount" xml:"TotalCount" require:"true"`
  CostUnitDtoList []*QueryCostUnitResponseDataCostUnitDtoList `json:"CostUnitDtoList" xml:"CostUnitDtoList" require:"true" type:"Repeated"`
}

func (s QueryCostUnitResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCostUnitResponseData) GoString() string {
  return s.String()
}

func (s *QueryCostUnitResponseData) SetPageNum(v int) *QueryCostUnitResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryCostUnitResponseData) SetPageSize(v int) *QueryCostUnitResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryCostUnitResponseData) SetTotalCount(v int) *QueryCostUnitResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryCostUnitResponseData) SetCostUnitDtoList(v []*QueryCostUnitResponseDataCostUnitDtoList) *QueryCostUnitResponseData {
  s.CostUnitDtoList = v
  return s
}

type QueryCostUnitResponseDataCostUnitDtoList struct     {
  OwnerUid *int64 `json:"OwnerUid" xml:"OwnerUid" require:"true"`
  ParentUnitId *int64 `json:"ParentUnitId" xml:"ParentUnitId" require:"true"`
  UnitId *int64 `json:"UnitId" xml:"UnitId" require:"true"`
  UnitName *string `json:"UnitName" xml:"UnitName" require:"true"`
}

func (s QueryCostUnitResponseDataCostUnitDtoList) String() string {
  return tea.Prettify(s)
}

func (s QueryCostUnitResponseDataCostUnitDtoList) GoString() string {
  return s.String()
}

func (s *QueryCostUnitResponseDataCostUnitDtoList) SetOwnerUid(v int64) *QueryCostUnitResponseDataCostUnitDtoList {
  s.OwnerUid = &v
  return s
}

func (s *QueryCostUnitResponseDataCostUnitDtoList) SetParentUnitId(v int64) *QueryCostUnitResponseDataCostUnitDtoList {
  s.ParentUnitId = &v
  return s
}

func (s *QueryCostUnitResponseDataCostUnitDtoList) SetUnitId(v int64) *QueryCostUnitResponseDataCostUnitDtoList {
  s.UnitId = &v
  return s
}

func (s *QueryCostUnitResponseDataCostUnitDtoList) SetUnitName(v string) *QueryCostUnitResponseDataCostUnitDtoList {
  s.UnitName = &v
  return s
}

type DeleteCostUnitRequest struct {
  OwnerUid *int64 `json:"OwnerUid" xml:"OwnerUid" require:"true"`
  UnitId *int64 `json:"UnitId" xml:"UnitId" require:"true"`
}

func (s DeleteCostUnitRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteCostUnitRequest) GoString() string {
  return s.String()
}

func (s *DeleteCostUnitRequest) SetOwnerUid(v int64) *DeleteCostUnitRequest {
  s.OwnerUid = &v
  return s
}

func (s *DeleteCostUnitRequest) SetUnitId(v int64) *DeleteCostUnitRequest {
  s.UnitId = &v
  return s
}

type DeleteCostUnitResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *DeleteCostUnitResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s DeleteCostUnitResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteCostUnitResponse) GoString() string {
  return s.String()
}

func (s *DeleteCostUnitResponse) SetRequestId(v string) *DeleteCostUnitResponse {
  s.RequestId = &v
  return s
}

func (s *DeleteCostUnitResponse) SetSuccess(v bool) *DeleteCostUnitResponse {
  s.Success = &v
  return s
}

func (s *DeleteCostUnitResponse) SetCode(v string) *DeleteCostUnitResponse {
  s.Code = &v
  return s
}

func (s *DeleteCostUnitResponse) SetMessage(v string) *DeleteCostUnitResponse {
  s.Message = &v
  return s
}

func (s *DeleteCostUnitResponse) SetData(v *DeleteCostUnitResponseData) *DeleteCostUnitResponse {
  s.Data = v
  return s
}

type DeleteCostUnitResponseData struct {
  OwnerUid *int64 `json:"OwnerUid" xml:"OwnerUid" require:"true"`
  UnitId *int64 `json:"UnitId" xml:"UnitId" require:"true"`
  IsSuccess *bool `json:"IsSuccess" xml:"IsSuccess" require:"true"`
}

func (s DeleteCostUnitResponseData) String() string {
  return tea.Prettify(s)
}

func (s DeleteCostUnitResponseData) GoString() string {
  return s.String()
}

func (s *DeleteCostUnitResponseData) SetOwnerUid(v int64) *DeleteCostUnitResponseData {
  s.OwnerUid = &v
  return s
}

func (s *DeleteCostUnitResponseData) SetUnitId(v int64) *DeleteCostUnitResponseData {
  s.UnitId = &v
  return s
}

func (s *DeleteCostUnitResponseData) SetIsSuccess(v bool) *DeleteCostUnitResponseData {
  s.IsSuccess = &v
  return s
}

type AllocateCostUnitResourceRequest struct {
  FromUnitUserId *int64 `json:"FromUnitUserId" xml:"FromUnitUserId" require:"true"`
  FromUnitId *int64 `json:"FromUnitId" xml:"FromUnitId" require:"true"`
  ResourceInstanceList []*AllocateCostUnitResourceRequestResourceInstanceList `json:"ResourceInstanceList" xml:"ResourceInstanceList" require:"true" type:"Repeated"`
  ToUnitUserId *int64 `json:"ToUnitUserId" xml:"ToUnitUserId" require:"true"`
  ToUnitId *int64 `json:"ToUnitId" xml:"ToUnitId" require:"true"`
}

func (s AllocateCostUnitResourceRequest) String() string {
  return tea.Prettify(s)
}

func (s AllocateCostUnitResourceRequest) GoString() string {
  return s.String()
}

func (s *AllocateCostUnitResourceRequest) SetFromUnitUserId(v int64) *AllocateCostUnitResourceRequest {
  s.FromUnitUserId = &v
  return s
}

func (s *AllocateCostUnitResourceRequest) SetFromUnitId(v int64) *AllocateCostUnitResourceRequest {
  s.FromUnitId = &v
  return s
}

func (s *AllocateCostUnitResourceRequest) SetResourceInstanceList(v []*AllocateCostUnitResourceRequestResourceInstanceList) *AllocateCostUnitResourceRequest {
  s.ResourceInstanceList = v
  return s
}

func (s *AllocateCostUnitResourceRequest) SetToUnitUserId(v int64) *AllocateCostUnitResourceRequest {
  s.ToUnitUserId = &v
  return s
}

func (s *AllocateCostUnitResourceRequest) SetToUnitId(v int64) *AllocateCostUnitResourceRequest {
  s.ToUnitId = &v
  return s
}

type AllocateCostUnitResourceRequestResourceInstanceList struct     {
  ResourceUserId *int64 `json:"ResourceUserId" xml:"ResourceUserId" require:"true"`
  ResourceId *string `json:"ResourceId" xml:"ResourceId" require:"true"`
  CommodityCode *string `json:"CommodityCode" xml:"CommodityCode" require:"true"`
}

func (s AllocateCostUnitResourceRequestResourceInstanceList) String() string {
  return tea.Prettify(s)
}

func (s AllocateCostUnitResourceRequestResourceInstanceList) GoString() string {
  return s.String()
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) SetResourceUserId(v int64) *AllocateCostUnitResourceRequestResourceInstanceList {
  s.ResourceUserId = &v
  return s
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) SetResourceId(v string) *AllocateCostUnitResourceRequestResourceInstanceList {
  s.ResourceId = &v
  return s
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) SetCommodityCode(v string) *AllocateCostUnitResourceRequestResourceInstanceList {
  s.CommodityCode = &v
  return s
}

type AllocateCostUnitResourceResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *AllocateCostUnitResourceResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s AllocateCostUnitResourceResponse) String() string {
  return tea.Prettify(s)
}

func (s AllocateCostUnitResourceResponse) GoString() string {
  return s.String()
}

func (s *AllocateCostUnitResourceResponse) SetRequestId(v string) *AllocateCostUnitResourceResponse {
  s.RequestId = &v
  return s
}

func (s *AllocateCostUnitResourceResponse) SetSuccess(v bool) *AllocateCostUnitResourceResponse {
  s.Success = &v
  return s
}

func (s *AllocateCostUnitResourceResponse) SetCode(v string) *AllocateCostUnitResourceResponse {
  s.Code = &v
  return s
}

func (s *AllocateCostUnitResourceResponse) SetMessage(v string) *AllocateCostUnitResourceResponse {
  s.Message = &v
  return s
}

func (s *AllocateCostUnitResourceResponse) SetData(v *AllocateCostUnitResourceResponseData) *AllocateCostUnitResourceResponse {
  s.Data = v
  return s
}

type AllocateCostUnitResourceResponseData struct {
  ToUnitUserId *int64 `json:"ToUnitUserId" xml:"ToUnitUserId" require:"true"`
  ToUnitId *int64 `json:"ToUnitId" xml:"ToUnitId" require:"true"`
  IsSuccess *bool `json:"IsSuccess" xml:"IsSuccess" require:"true"`
}

func (s AllocateCostUnitResourceResponseData) String() string {
  return tea.Prettify(s)
}

func (s AllocateCostUnitResourceResponseData) GoString() string {
  return s.String()
}

func (s *AllocateCostUnitResourceResponseData) SetToUnitUserId(v int64) *AllocateCostUnitResourceResponseData {
  s.ToUnitUserId = &v
  return s
}

func (s *AllocateCostUnitResourceResponseData) SetToUnitId(v int64) *AllocateCostUnitResourceResponseData {
  s.ToUnitId = &v
  return s
}

func (s *AllocateCostUnitResourceResponseData) SetIsSuccess(v bool) *AllocateCostUnitResourceResponseData {
  s.IsSuccess = &v
  return s
}

type QueryCostUnitResourceRequest struct {
  OwnerUid *int64 `json:"OwnerUid" xml:"OwnerUid" require:"true"`
  UnitId *int64 `json:"UnitId" xml:"UnitId" require:"true"`
  PageNum *int `json:"PageNum" xml:"PageNum"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
}

func (s QueryCostUnitResourceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCostUnitResourceRequest) GoString() string {
  return s.String()
}

func (s *QueryCostUnitResourceRequest) SetOwnerUid(v int64) *QueryCostUnitResourceRequest {
  s.OwnerUid = &v
  return s
}

func (s *QueryCostUnitResourceRequest) SetUnitId(v int64) *QueryCostUnitResourceRequest {
  s.UnitId = &v
  return s
}

func (s *QueryCostUnitResourceRequest) SetPageNum(v int) *QueryCostUnitResourceRequest {
  s.PageNum = &v
  return s
}

func (s *QueryCostUnitResourceRequest) SetPageSize(v int) *QueryCostUnitResourceRequest {
  s.PageSize = &v
  return s
}

type QueryCostUnitResourceResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryCostUnitResourceResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryCostUnitResourceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCostUnitResourceResponse) GoString() string {
  return s.String()
}

func (s *QueryCostUnitResourceResponse) SetRequestId(v string) *QueryCostUnitResourceResponse {
  s.RequestId = &v
  return s
}

func (s *QueryCostUnitResourceResponse) SetSuccess(v bool) *QueryCostUnitResourceResponse {
  s.Success = &v
  return s
}

func (s *QueryCostUnitResourceResponse) SetCode(v string) *QueryCostUnitResourceResponse {
  s.Code = &v
  return s
}

func (s *QueryCostUnitResourceResponse) SetMessage(v string) *QueryCostUnitResourceResponse {
  s.Message = &v
  return s
}

func (s *QueryCostUnitResourceResponse) SetData(v *QueryCostUnitResourceResponseData) *QueryCostUnitResourceResponse {
  s.Data = v
  return s
}

type QueryCostUnitResourceResponseData struct {
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  TotalCount *int `json:"TotalCount" xml:"TotalCount" require:"true"`
  ResourceInstanceDtoList []*QueryCostUnitResourceResponseDataResourceInstanceDtoList `json:"ResourceInstanceDtoList" xml:"ResourceInstanceDtoList" require:"true" type:"Repeated"`
  CostUnit *QueryCostUnitResourceResponseDataCostUnit `json:"CostUnit" xml:"CostUnit" require:"true" type:"Struct"`
  CostUnitStatisInfo *QueryCostUnitResourceResponseDataCostUnitStatisInfo `json:"CostUnitStatisInfo" xml:"CostUnitStatisInfo" require:"true" type:"Struct"`
}

func (s QueryCostUnitResourceResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCostUnitResourceResponseData) GoString() string {
  return s.String()
}

func (s *QueryCostUnitResourceResponseData) SetPageNum(v int) *QueryCostUnitResourceResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryCostUnitResourceResponseData) SetPageSize(v int) *QueryCostUnitResourceResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryCostUnitResourceResponseData) SetTotalCount(v int) *QueryCostUnitResourceResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryCostUnitResourceResponseData) SetResourceInstanceDtoList(v []*QueryCostUnitResourceResponseDataResourceInstanceDtoList) *QueryCostUnitResourceResponseData {
  s.ResourceInstanceDtoList = v
  return s
}

func (s *QueryCostUnitResourceResponseData) SetCostUnit(v *QueryCostUnitResourceResponseDataCostUnit) *QueryCostUnitResourceResponseData {
  s.CostUnit = v
  return s
}

func (s *QueryCostUnitResourceResponseData) SetCostUnitStatisInfo(v *QueryCostUnitResourceResponseDataCostUnitStatisInfo) *QueryCostUnitResourceResponseData {
  s.CostUnitStatisInfo = v
  return s
}

type QueryCostUnitResourceResponseDataResourceInstanceDtoList struct     {
  ResourceUserId *int64 `json:"ResourceUserId" xml:"ResourceUserId" require:"true"`
  ResourceId *string `json:"ResourceId" xml:"ResourceId" require:"true"`
  CommodityCode *string `json:"CommodityCode" xml:"CommodityCode" require:"true"`
  ResourceUserName *string `json:"ResourceUserName" xml:"ResourceUserName" require:"true"`
  CommodityName *string `json:"CommodityName" xml:"CommodityName" require:"true"`
  ResourceGroup *string `json:"ResourceGroup" xml:"ResourceGroup" require:"true"`
  ResourceTag *string `json:"ResourceTag" xml:"ResourceTag" require:"true"`
  ResourceNick *string `json:"ResourceNick" xml:"ResourceNick" require:"true"`
  ResourceType *string `json:"ResourceType" xml:"ResourceType" require:"true"`
  ResourceStatus *string `json:"ResourceStatus" xml:"ResourceStatus" require:"true"`
  RelatedResources *string `json:"RelatedResources" xml:"RelatedResources" require:"true"`
  ApportionCode *string `json:"ApportionCode" xml:"ApportionCode" require:"true"`
  ApportionName *string `json:"ApportionName" xml:"ApportionName" require:"true"`
}

func (s QueryCostUnitResourceResponseDataResourceInstanceDtoList) String() string {
  return tea.Prettify(s)
}

func (s QueryCostUnitResourceResponseDataResourceInstanceDtoList) GoString() string {
  return s.String()
}

func (s *QueryCostUnitResourceResponseDataResourceInstanceDtoList) SetResourceUserId(v int64) *QueryCostUnitResourceResponseDataResourceInstanceDtoList {
  s.ResourceUserId = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataResourceInstanceDtoList) SetResourceId(v string) *QueryCostUnitResourceResponseDataResourceInstanceDtoList {
  s.ResourceId = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataResourceInstanceDtoList) SetCommodityCode(v string) *QueryCostUnitResourceResponseDataResourceInstanceDtoList {
  s.CommodityCode = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataResourceInstanceDtoList) SetResourceUserName(v string) *QueryCostUnitResourceResponseDataResourceInstanceDtoList {
  s.ResourceUserName = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataResourceInstanceDtoList) SetCommodityName(v string) *QueryCostUnitResourceResponseDataResourceInstanceDtoList {
  s.CommodityName = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataResourceInstanceDtoList) SetResourceGroup(v string) *QueryCostUnitResourceResponseDataResourceInstanceDtoList {
  s.ResourceGroup = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataResourceInstanceDtoList) SetResourceTag(v string) *QueryCostUnitResourceResponseDataResourceInstanceDtoList {
  s.ResourceTag = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataResourceInstanceDtoList) SetResourceNick(v string) *QueryCostUnitResourceResponseDataResourceInstanceDtoList {
  s.ResourceNick = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataResourceInstanceDtoList) SetResourceType(v string) *QueryCostUnitResourceResponseDataResourceInstanceDtoList {
  s.ResourceType = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataResourceInstanceDtoList) SetResourceStatus(v string) *QueryCostUnitResourceResponseDataResourceInstanceDtoList {
  s.ResourceStatus = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataResourceInstanceDtoList) SetRelatedResources(v string) *QueryCostUnitResourceResponseDataResourceInstanceDtoList {
  s.RelatedResources = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataResourceInstanceDtoList) SetApportionCode(v string) *QueryCostUnitResourceResponseDataResourceInstanceDtoList {
  s.ApportionCode = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataResourceInstanceDtoList) SetApportionName(v string) *QueryCostUnitResourceResponseDataResourceInstanceDtoList {
  s.ApportionName = &v
  return s
}

type QueryCostUnitResourceResponseDataCostUnit struct {
  OwnerUid *int64 `json:"OwnerUid" xml:"OwnerUid" require:"true"`
  ParentUnitId *int64 `json:"ParentUnitId" xml:"ParentUnitId" require:"true"`
  UnitId *int64 `json:"UnitId" xml:"UnitId" require:"true"`
  UnitName *string `json:"UnitName" xml:"UnitName" require:"true"`
}

func (s QueryCostUnitResourceResponseDataCostUnit) String() string {
  return tea.Prettify(s)
}

func (s QueryCostUnitResourceResponseDataCostUnit) GoString() string {
  return s.String()
}

func (s *QueryCostUnitResourceResponseDataCostUnit) SetOwnerUid(v int64) *QueryCostUnitResourceResponseDataCostUnit {
  s.OwnerUid = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataCostUnit) SetParentUnitId(v int64) *QueryCostUnitResourceResponseDataCostUnit {
  s.ParentUnitId = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataCostUnit) SetUnitId(v int64) *QueryCostUnitResourceResponseDataCostUnit {
  s.UnitId = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataCostUnit) SetUnitName(v string) *QueryCostUnitResourceResponseDataCostUnit {
  s.UnitName = &v
  return s
}

type QueryCostUnitResourceResponseDataCostUnitStatisInfo struct {
  ResourceCount *int64 `json:"ResourceCount" xml:"ResourceCount" require:"true"`
  ResourceGroupCount *int64 `json:"ResourceGroupCount" xml:"ResourceGroupCount" require:"true"`
  SubUnitCount *int64 `json:"SubUnitCount" xml:"SubUnitCount" require:"true"`
  UserCount *int64 `json:"UserCount" xml:"UserCount" require:"true"`
  TotalResourceCount *int64 `json:"TotalResourceCount" xml:"TotalResourceCount" require:"true"`
  TotalUserCount *int64 `json:"TotalUserCount" xml:"TotalUserCount" require:"true"`
  TotalResourceGroupCount *int64 `json:"TotalResourceGroupCount" xml:"TotalResourceGroupCount" require:"true"`
}

func (s QueryCostUnitResourceResponseDataCostUnitStatisInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryCostUnitResourceResponseDataCostUnitStatisInfo) GoString() string {
  return s.String()
}

func (s *QueryCostUnitResourceResponseDataCostUnitStatisInfo) SetResourceCount(v int64) *QueryCostUnitResourceResponseDataCostUnitStatisInfo {
  s.ResourceCount = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataCostUnitStatisInfo) SetResourceGroupCount(v int64) *QueryCostUnitResourceResponseDataCostUnitStatisInfo {
  s.ResourceGroupCount = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataCostUnitStatisInfo) SetSubUnitCount(v int64) *QueryCostUnitResourceResponseDataCostUnitStatisInfo {
  s.SubUnitCount = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataCostUnitStatisInfo) SetUserCount(v int64) *QueryCostUnitResourceResponseDataCostUnitStatisInfo {
  s.UserCount = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataCostUnitStatisInfo) SetTotalResourceCount(v int64) *QueryCostUnitResourceResponseDataCostUnitStatisInfo {
  s.TotalResourceCount = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataCostUnitStatisInfo) SetTotalUserCount(v int64) *QueryCostUnitResourceResponseDataCostUnitStatisInfo {
  s.TotalUserCount = &v
  return s
}

func (s *QueryCostUnitResourceResponseDataCostUnitStatisInfo) SetTotalResourceGroupCount(v int64) *QueryCostUnitResourceResponseDataCostUnitStatisInfo {
  s.TotalResourceGroupCount = &v
  return s
}

type RenewResourcePackageRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  InstanceId *string `json:"InstanceId" xml:"InstanceId" require:"true"`
  EffectiveDate *string `json:"EffectiveDate" xml:"EffectiveDate"`
  Duration *int `json:"Duration" xml:"Duration" require:"true"`
  PricingCycle *string `json:"PricingCycle" xml:"PricingCycle" require:"true"`
}

func (s RenewResourcePackageRequest) String() string {
  return tea.Prettify(s)
}

func (s RenewResourcePackageRequest) GoString() string {
  return s.String()
}

func (s *RenewResourcePackageRequest) SetOwnerId(v int64) *RenewResourcePackageRequest {
  s.OwnerId = &v
  return s
}

func (s *RenewResourcePackageRequest) SetInstanceId(v string) *RenewResourcePackageRequest {
  s.InstanceId = &v
  return s
}

func (s *RenewResourcePackageRequest) SetEffectiveDate(v string) *RenewResourcePackageRequest {
  s.EffectiveDate = &v
  return s
}

func (s *RenewResourcePackageRequest) SetDuration(v int) *RenewResourcePackageRequest {
  s.Duration = &v
  return s
}

func (s *RenewResourcePackageRequest) SetPricingCycle(v string) *RenewResourcePackageRequest {
  s.PricingCycle = &v
  return s
}

type RenewResourcePackageResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  OrderId *int64 `json:"OrderId" xml:"OrderId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *RenewResourcePackageResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s RenewResourcePackageResponse) String() string {
  return tea.Prettify(s)
}

func (s RenewResourcePackageResponse) GoString() string {
  return s.String()
}

func (s *RenewResourcePackageResponse) SetRequestId(v string) *RenewResourcePackageResponse {
  s.RequestId = &v
  return s
}

func (s *RenewResourcePackageResponse) SetOrderId(v int64) *RenewResourcePackageResponse {
  s.OrderId = &v
  return s
}

func (s *RenewResourcePackageResponse) SetSuccess(v bool) *RenewResourcePackageResponse {
  s.Success = &v
  return s
}

func (s *RenewResourcePackageResponse) SetCode(v string) *RenewResourcePackageResponse {
  s.Code = &v
  return s
}

func (s *RenewResourcePackageResponse) SetMessage(v string) *RenewResourcePackageResponse {
  s.Message = &v
  return s
}

func (s *RenewResourcePackageResponse) SetData(v *RenewResourcePackageResponseData) *RenewResourcePackageResponse {
  s.Data = v
  return s
}

type RenewResourcePackageResponseData struct {
  OrderId *int64 `json:"OrderId" xml:"OrderId" require:"true"`
  InstanceId *string `json:"InstanceId" xml:"InstanceId" require:"true"`
}

func (s RenewResourcePackageResponseData) String() string {
  return tea.Prettify(s)
}

func (s RenewResourcePackageResponseData) GoString() string {
  return s.String()
}

func (s *RenewResourcePackageResponseData) SetOrderId(v int64) *RenewResourcePackageResponseData {
  s.OrderId = &v
  return s
}

func (s *RenewResourcePackageResponseData) SetInstanceId(v string) *RenewResourcePackageResponseData {
  s.InstanceId = &v
  return s
}

type UpgradeResourcePackageRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  InstanceId *string `json:"InstanceId" xml:"InstanceId"`
  EffectiveDate *string `json:"EffectiveDate" xml:"EffectiveDate"`
  Specification *string `json:"Specification" xml:"Specification"`
}

func (s UpgradeResourcePackageRequest) String() string {
  return tea.Prettify(s)
}

func (s UpgradeResourcePackageRequest) GoString() string {
  return s.String()
}

func (s *UpgradeResourcePackageRequest) SetOwnerId(v int64) *UpgradeResourcePackageRequest {
  s.OwnerId = &v
  return s
}

func (s *UpgradeResourcePackageRequest) SetInstanceId(v string) *UpgradeResourcePackageRequest {
  s.InstanceId = &v
  return s
}

func (s *UpgradeResourcePackageRequest) SetEffectiveDate(v string) *UpgradeResourcePackageRequest {
  s.EffectiveDate = &v
  return s
}

func (s *UpgradeResourcePackageRequest) SetSpecification(v string) *UpgradeResourcePackageRequest {
  s.Specification = &v
  return s
}

type UpgradeResourcePackageResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  OrderId *int64 `json:"OrderId" xml:"OrderId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *UpgradeResourcePackageResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s UpgradeResourcePackageResponse) String() string {
  return tea.Prettify(s)
}

func (s UpgradeResourcePackageResponse) GoString() string {
  return s.String()
}

func (s *UpgradeResourcePackageResponse) SetRequestId(v string) *UpgradeResourcePackageResponse {
  s.RequestId = &v
  return s
}

func (s *UpgradeResourcePackageResponse) SetOrderId(v int64) *UpgradeResourcePackageResponse {
  s.OrderId = &v
  return s
}

func (s *UpgradeResourcePackageResponse) SetSuccess(v bool) *UpgradeResourcePackageResponse {
  s.Success = &v
  return s
}

func (s *UpgradeResourcePackageResponse) SetCode(v string) *UpgradeResourcePackageResponse {
  s.Code = &v
  return s
}

func (s *UpgradeResourcePackageResponse) SetMessage(v string) *UpgradeResourcePackageResponse {
  s.Message = &v
  return s
}

func (s *UpgradeResourcePackageResponse) SetData(v *UpgradeResourcePackageResponseData) *UpgradeResourcePackageResponse {
  s.Data = v
  return s
}

type UpgradeResourcePackageResponseData struct {
  OrderId *int64 `json:"OrderId" xml:"OrderId" require:"true"`
  InstanceId *string `json:"InstanceId" xml:"InstanceId" require:"true"`
}

func (s UpgradeResourcePackageResponseData) String() string {
  return tea.Prettify(s)
}

func (s UpgradeResourcePackageResponseData) GoString() string {
  return s.String()
}

func (s *UpgradeResourcePackageResponseData) SetOrderId(v int64) *UpgradeResourcePackageResponseData {
  s.OrderId = &v
  return s
}

func (s *UpgradeResourcePackageResponseData) SetInstanceId(v string) *UpgradeResourcePackageResponseData {
  s.InstanceId = &v
  return s
}

type CreateAgAccountRequest struct {
  LoginEmail *string `json:"LoginEmail" xml:"LoginEmail" require:"true"`
  AccountAttr *string `json:"AccountAttr" xml:"AccountAttr"`
  EnterpriseName *string `json:"EnterpriseName" xml:"EnterpriseName"`
  FirstName *string `json:"FirstName" xml:"FirstName"`
  LastName *string `json:"LastName" xml:"LastName"`
  NationCode *string `json:"NationCode" xml:"NationCode"`
  ProvinceName *string `json:"ProvinceName" xml:"ProvinceName"`
  CityName *string `json:"CityName" xml:"CityName"`
  Postcode *string `json:"Postcode" xml:"Postcode"`
}

func (s CreateAgAccountRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateAgAccountRequest) GoString() string {
  return s.String()
}

func (s *CreateAgAccountRequest) SetLoginEmail(v string) *CreateAgAccountRequest {
  s.LoginEmail = &v
  return s
}

func (s *CreateAgAccountRequest) SetAccountAttr(v string) *CreateAgAccountRequest {
  s.AccountAttr = &v
  return s
}

func (s *CreateAgAccountRequest) SetEnterpriseName(v string) *CreateAgAccountRequest {
  s.EnterpriseName = &v
  return s
}

func (s *CreateAgAccountRequest) SetFirstName(v string) *CreateAgAccountRequest {
  s.FirstName = &v
  return s
}

func (s *CreateAgAccountRequest) SetLastName(v string) *CreateAgAccountRequest {
  s.LastName = &v
  return s
}

func (s *CreateAgAccountRequest) SetNationCode(v string) *CreateAgAccountRequest {
  s.NationCode = &v
  return s
}

func (s *CreateAgAccountRequest) SetProvinceName(v string) *CreateAgAccountRequest {
  s.ProvinceName = &v
  return s
}

func (s *CreateAgAccountRequest) SetCityName(v string) *CreateAgAccountRequest {
  s.CityName = &v
  return s
}

func (s *CreateAgAccountRequest) SetPostcode(v string) *CreateAgAccountRequest {
  s.Postcode = &v
  return s
}

type CreateAgAccountResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  AgRelationDto *CreateAgAccountResponseAgRelationDto `json:"AgRelationDto" xml:"AgRelationDto" require:"true" type:"Struct"`
}

func (s CreateAgAccountResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateAgAccountResponse) GoString() string {
  return s.String()
}

func (s *CreateAgAccountResponse) SetRequestId(v string) *CreateAgAccountResponse {
  s.RequestId = &v
  return s
}

func (s *CreateAgAccountResponse) SetCode(v string) *CreateAgAccountResponse {
  s.Code = &v
  return s
}

func (s *CreateAgAccountResponse) SetMessage(v string) *CreateAgAccountResponse {
  s.Message = &v
  return s
}

func (s *CreateAgAccountResponse) SetSuccess(v bool) *CreateAgAccountResponse {
  s.Success = &v
  return s
}

func (s *CreateAgAccountResponse) SetAgRelationDto(v *CreateAgAccountResponseAgRelationDto) *CreateAgAccountResponse {
  s.AgRelationDto = v
  return s
}

type CreateAgAccountResponseAgRelationDto struct {
  Pk *string `json:"Pk" xml:"Pk" require:"true"`
  Type *string `json:"Type" xml:"Type" require:"true"`
  Mpk *string `json:"Mpk" xml:"Mpk" require:"true"`
  RamAdminRoleName *string `json:"RamAdminRoleName" xml:"RamAdminRoleName" require:"true"`
}

func (s CreateAgAccountResponseAgRelationDto) String() string {
  return tea.Prettify(s)
}

func (s CreateAgAccountResponseAgRelationDto) GoString() string {
  return s.String()
}

func (s *CreateAgAccountResponseAgRelationDto) SetPk(v string) *CreateAgAccountResponseAgRelationDto {
  s.Pk = &v
  return s
}

func (s *CreateAgAccountResponseAgRelationDto) SetType(v string) *CreateAgAccountResponseAgRelationDto {
  s.Type = &v
  return s
}

func (s *CreateAgAccountResponseAgRelationDto) SetMpk(v string) *CreateAgAccountResponseAgRelationDto {
  s.Mpk = &v
  return s
}

func (s *CreateAgAccountResponseAgRelationDto) SetRamAdminRoleName(v string) *CreateAgAccountResponseAgRelationDto {
  s.RamAdminRoleName = &v
  return s
}

type GetCustomerAccountInfoRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId" require:"true"`
}

func (s GetCustomerAccountInfoRequest) String() string {
  return tea.Prettify(s)
}

func (s GetCustomerAccountInfoRequest) GoString() string {
  return s.String()
}

func (s *GetCustomerAccountInfoRequest) SetOwnerId(v int64) *GetCustomerAccountInfoRequest {
  s.OwnerId = &v
  return s
}

type GetCustomerAccountInfoResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *GetCustomerAccountInfoResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s GetCustomerAccountInfoResponse) String() string {
  return tea.Prettify(s)
}

func (s GetCustomerAccountInfoResponse) GoString() string {
  return s.String()
}

func (s *GetCustomerAccountInfoResponse) SetRequestId(v string) *GetCustomerAccountInfoResponse {
  s.RequestId = &v
  return s
}

func (s *GetCustomerAccountInfoResponse) SetSuccess(v bool) *GetCustomerAccountInfoResponse {
  s.Success = &v
  return s
}

func (s *GetCustomerAccountInfoResponse) SetCode(v string) *GetCustomerAccountInfoResponse {
  s.Code = &v
  return s
}

func (s *GetCustomerAccountInfoResponse) SetMessage(v string) *GetCustomerAccountInfoResponse {
  s.Message = &v
  return s
}

func (s *GetCustomerAccountInfoResponse) SetData(v *GetCustomerAccountInfoResponseData) *GetCustomerAccountInfoResponse {
  s.Data = v
  return s
}

type GetCustomerAccountInfoResponseData struct {
  LoginEmail *string `json:"LoginEmail" xml:"LoginEmail" require:"true"`
  AccountType *string `json:"AccountType" xml:"AccountType" require:"true"`
  Mpk *int64 `json:"Mpk" xml:"Mpk" require:"true"`
  HostingStatus *string `json:"HostingStatus" xml:"HostingStatus" require:"true"`
  CreditLimitStatus *string `json:"CreditLimitStatus" xml:"CreditLimitStatus" require:"true"`
  IsCertified *bool `json:"IsCertified" xml:"IsCertified" require:"true"`
}

func (s GetCustomerAccountInfoResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetCustomerAccountInfoResponseData) GoString() string {
  return s.String()
}

func (s *GetCustomerAccountInfoResponseData) SetLoginEmail(v string) *GetCustomerAccountInfoResponseData {
  s.LoginEmail = &v
  return s
}

func (s *GetCustomerAccountInfoResponseData) SetAccountType(v string) *GetCustomerAccountInfoResponseData {
  s.AccountType = &v
  return s
}

func (s *GetCustomerAccountInfoResponseData) SetMpk(v int64) *GetCustomerAccountInfoResponseData {
  s.Mpk = &v
  return s
}

func (s *GetCustomerAccountInfoResponseData) SetHostingStatus(v string) *GetCustomerAccountInfoResponseData {
  s.HostingStatus = &v
  return s
}

func (s *GetCustomerAccountInfoResponseData) SetCreditLimitStatus(v string) *GetCustomerAccountInfoResponseData {
  s.CreditLimitStatus = &v
  return s
}

func (s *GetCustomerAccountInfoResponseData) SetIsCertified(v bool) *GetCustomerAccountInfoResponseData {
  s.IsCertified = &v
  return s
}

type GetCustomerListRequest struct {
}

func (s GetCustomerListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetCustomerListRequest) GoString() string {
  return s.String()
}

type GetCustomerListResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *GetCustomerListResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s GetCustomerListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetCustomerListResponse) GoString() string {
  return s.String()
}

func (s *GetCustomerListResponse) SetRequestId(v string) *GetCustomerListResponse {
  s.RequestId = &v
  return s
}

func (s *GetCustomerListResponse) SetSuccess(v bool) *GetCustomerListResponse {
  s.Success = &v
  return s
}

func (s *GetCustomerListResponse) SetCode(v string) *GetCustomerListResponse {
  s.Code = &v
  return s
}

func (s *GetCustomerListResponse) SetMessage(v string) *GetCustomerListResponse {
  s.Message = &v
  return s
}

func (s *GetCustomerListResponse) SetData(v *GetCustomerListResponseData) *GetCustomerListResponse {
  s.Data = v
  return s
}

type GetCustomerListResponseData struct {
  UidList []*string `json:"UidList" xml:"UidList" require:"true" type:"Repeated"`
}

func (s GetCustomerListResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetCustomerListResponseData) GoString() string {
  return s.String()
}

func (s *GetCustomerListResponseData) SetUidList(v []*string) *GetCustomerListResponseData {
  s.UidList = v
  return s
}

type ChangeResellerConsumeAmountRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId" require:"true"`
  AdjustType *string `json:"AdjustType" xml:"AdjustType" require:"true"`
  Amount *string `json:"Amount" xml:"Amount" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  BusinessType *string `json:"BusinessType" xml:"BusinessType" require:"true"`
  Source *string `json:"Source" xml:"Source" require:"true"`
  OutBizId *string `json:"OutBizId" xml:"OutBizId" require:"true"`
  ExtendMap *string `json:"ExtendMap" xml:"ExtendMap"`
}

func (s ChangeResellerConsumeAmountRequest) String() string {
  return tea.Prettify(s)
}

func (s ChangeResellerConsumeAmountRequest) GoString() string {
  return s.String()
}

func (s *ChangeResellerConsumeAmountRequest) SetOwnerId(v int64) *ChangeResellerConsumeAmountRequest {
  s.OwnerId = &v
  return s
}

func (s *ChangeResellerConsumeAmountRequest) SetAdjustType(v string) *ChangeResellerConsumeAmountRequest {
  s.AdjustType = &v
  return s
}

func (s *ChangeResellerConsumeAmountRequest) SetAmount(v string) *ChangeResellerConsumeAmountRequest {
  s.Amount = &v
  return s
}

func (s *ChangeResellerConsumeAmountRequest) SetCurrency(v string) *ChangeResellerConsumeAmountRequest {
  s.Currency = &v
  return s
}

func (s *ChangeResellerConsumeAmountRequest) SetBusinessType(v string) *ChangeResellerConsumeAmountRequest {
  s.BusinessType = &v
  return s
}

func (s *ChangeResellerConsumeAmountRequest) SetSource(v string) *ChangeResellerConsumeAmountRequest {
  s.Source = &v
  return s
}

func (s *ChangeResellerConsumeAmountRequest) SetOutBizId(v string) *ChangeResellerConsumeAmountRequest {
  s.OutBizId = &v
  return s
}

func (s *ChangeResellerConsumeAmountRequest) SetExtendMap(v string) *ChangeResellerConsumeAmountRequest {
  s.ExtendMap = &v
  return s
}

type ChangeResellerConsumeAmountResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Data *string `json:"Data" xml:"Data" require:"true"`
}

func (s ChangeResellerConsumeAmountResponse) String() string {
  return tea.Prettify(s)
}

func (s ChangeResellerConsumeAmountResponse) GoString() string {
  return s.String()
}

func (s *ChangeResellerConsumeAmountResponse) SetRequestId(v string) *ChangeResellerConsumeAmountResponse {
  s.RequestId = &v
  return s
}

func (s *ChangeResellerConsumeAmountResponse) SetCode(v string) *ChangeResellerConsumeAmountResponse {
  s.Code = &v
  return s
}

func (s *ChangeResellerConsumeAmountResponse) SetMessage(v string) *ChangeResellerConsumeAmountResponse {
  s.Message = &v
  return s
}

func (s *ChangeResellerConsumeAmountResponse) SetSuccess(v bool) *ChangeResellerConsumeAmountResponse {
  s.Success = &v
  return s
}

func (s *ChangeResellerConsumeAmountResponse) SetData(v string) *ChangeResellerConsumeAmountResponse {
  s.Data = &v
  return s
}

type SetResellerUserStatusRequest struct {
  OwnerId *string `json:"OwnerId" xml:"OwnerId" require:"true"`
  Status *string `json:"Status" xml:"Status" require:"true"`
  BusinessType *string `json:"BusinessType" xml:"BusinessType" require:"true"`
}

func (s SetResellerUserStatusRequest) String() string {
  return tea.Prettify(s)
}

func (s SetResellerUserStatusRequest) GoString() string {
  return s.String()
}

func (s *SetResellerUserStatusRequest) SetOwnerId(v string) *SetResellerUserStatusRequest {
  s.OwnerId = &v
  return s
}

func (s *SetResellerUserStatusRequest) SetStatus(v string) *SetResellerUserStatusRequest {
  s.Status = &v
  return s
}

func (s *SetResellerUserStatusRequest) SetBusinessType(v string) *SetResellerUserStatusRequest {
  s.BusinessType = &v
  return s
}

type SetResellerUserStatusResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Data *bool `json:"Data" xml:"Data" require:"true"`
}

func (s SetResellerUserStatusResponse) String() string {
  return tea.Prettify(s)
}

func (s SetResellerUserStatusResponse) GoString() string {
  return s.String()
}

func (s *SetResellerUserStatusResponse) SetRequestId(v string) *SetResellerUserStatusResponse {
  s.RequestId = &v
  return s
}

func (s *SetResellerUserStatusResponse) SetCode(v string) *SetResellerUserStatusResponse {
  s.Code = &v
  return s
}

func (s *SetResellerUserStatusResponse) SetMessage(v string) *SetResellerUserStatusResponse {
  s.Message = &v
  return s
}

func (s *SetResellerUserStatusResponse) SetSuccess(v bool) *SetResellerUserStatusResponse {
  s.Success = &v
  return s
}

func (s *SetResellerUserStatusResponse) SetData(v bool) *SetResellerUserStatusResponse {
  s.Data = &v
  return s
}

type CreateResellerUserQuotaRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId" require:"true"`
  Amount *string `json:"Amount" xml:"Amount" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  OutBizId *string `json:"OutBizId" xml:"OutBizId"`
}

func (s CreateResellerUserQuotaRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateResellerUserQuotaRequest) GoString() string {
  return s.String()
}

func (s *CreateResellerUserQuotaRequest) SetOwnerId(v int64) *CreateResellerUserQuotaRequest {
  s.OwnerId = &v
  return s
}

func (s *CreateResellerUserQuotaRequest) SetAmount(v string) *CreateResellerUserQuotaRequest {
  s.Amount = &v
  return s
}

func (s *CreateResellerUserQuotaRequest) SetCurrency(v string) *CreateResellerUserQuotaRequest {
  s.Currency = &v
  return s
}

func (s *CreateResellerUserQuotaRequest) SetOutBizId(v string) *CreateResellerUserQuotaRequest {
  s.OutBizId = &v
  return s
}

type CreateResellerUserQuotaResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Data *bool `json:"Data" xml:"Data" require:"true"`
}

func (s CreateResellerUserQuotaResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateResellerUserQuotaResponse) GoString() string {
  return s.String()
}

func (s *CreateResellerUserQuotaResponse) SetRequestId(v string) *CreateResellerUserQuotaResponse {
  s.RequestId = &v
  return s
}

func (s *CreateResellerUserQuotaResponse) SetCode(v string) *CreateResellerUserQuotaResponse {
  s.Code = &v
  return s
}

func (s *CreateResellerUserQuotaResponse) SetMessage(v string) *CreateResellerUserQuotaResponse {
  s.Message = &v
  return s
}

func (s *CreateResellerUserQuotaResponse) SetSuccess(v bool) *CreateResellerUserQuotaResponse {
  s.Success = &v
  return s
}

func (s *CreateResellerUserQuotaResponse) SetData(v bool) *CreateResellerUserQuotaResponse {
  s.Data = &v
  return s
}

type SetResellerUserQuotaRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId" require:"true"`
  Amount *string `json:"Amount" xml:"Amount" require:"true"`
  Currency *string `json:"Currency" xml:"Currency"`
  OutBizId *string `json:"OutBizId" xml:"OutBizId"`
}

func (s SetResellerUserQuotaRequest) String() string {
  return tea.Prettify(s)
}

func (s SetResellerUserQuotaRequest) GoString() string {
  return s.String()
}

func (s *SetResellerUserQuotaRequest) SetOwnerId(v int64) *SetResellerUserQuotaRequest {
  s.OwnerId = &v
  return s
}

func (s *SetResellerUserQuotaRequest) SetAmount(v string) *SetResellerUserQuotaRequest {
  s.Amount = &v
  return s
}

func (s *SetResellerUserQuotaRequest) SetCurrency(v string) *SetResellerUserQuotaRequest {
  s.Currency = &v
  return s
}

func (s *SetResellerUserQuotaRequest) SetOutBizId(v string) *SetResellerUserQuotaRequest {
  s.OutBizId = &v
  return s
}

type SetResellerUserQuotaResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Data *bool `json:"Data" xml:"Data" require:"true"`
}

func (s SetResellerUserQuotaResponse) String() string {
  return tea.Prettify(s)
}

func (s SetResellerUserQuotaResponse) GoString() string {
  return s.String()
}

func (s *SetResellerUserQuotaResponse) SetRequestId(v string) *SetResellerUserQuotaResponse {
  s.RequestId = &v
  return s
}

func (s *SetResellerUserQuotaResponse) SetCode(v string) *SetResellerUserQuotaResponse {
  s.Code = &v
  return s
}

func (s *SetResellerUserQuotaResponse) SetMessage(v string) *SetResellerUserQuotaResponse {
  s.Message = &v
  return s
}

func (s *SetResellerUserQuotaResponse) SetSuccess(v bool) *SetResellerUserQuotaResponse {
  s.Success = &v
  return s
}

func (s *SetResellerUserQuotaResponse) SetData(v bool) *SetResellerUserQuotaResponse {
  s.Data = &v
  return s
}

type QueryResellerAvailableQuotaRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId" require:"true"`
  ItemCodes *string `json:"ItemCodes" xml:"ItemCodes"`
}

func (s QueryResellerAvailableQuotaRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryResellerAvailableQuotaRequest) GoString() string {
  return s.String()
}

func (s *QueryResellerAvailableQuotaRequest) SetOwnerId(v int64) *QueryResellerAvailableQuotaRequest {
  s.OwnerId = &v
  return s
}

func (s *QueryResellerAvailableQuotaRequest) SetItemCodes(v string) *QueryResellerAvailableQuotaRequest {
  s.ItemCodes = &v
  return s
}

type QueryResellerAvailableQuotaResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Data *string `json:"Data" xml:"Data" require:"true"`
}

func (s QueryResellerAvailableQuotaResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryResellerAvailableQuotaResponse) GoString() string {
  return s.String()
}

func (s *QueryResellerAvailableQuotaResponse) SetRequestId(v string) *QueryResellerAvailableQuotaResponse {
  s.RequestId = &v
  return s
}

func (s *QueryResellerAvailableQuotaResponse) SetCode(v string) *QueryResellerAvailableQuotaResponse {
  s.Code = &v
  return s
}

func (s *QueryResellerAvailableQuotaResponse) SetMessage(v string) *QueryResellerAvailableQuotaResponse {
  s.Message = &v
  return s
}

func (s *QueryResellerAvailableQuotaResponse) SetSuccess(v bool) *QueryResellerAvailableQuotaResponse {
  s.Success = &v
  return s
}

func (s *QueryResellerAvailableQuotaResponse) SetData(v string) *QueryResellerAvailableQuotaResponse {
  s.Data = &v
  return s
}

type SetResellerUserAlarmThresholdRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId" require:"true"`
  AlarmType *string `json:"AlarmType" xml:"AlarmType" require:"true"`
  AlarmThresholds *string `json:"AlarmThresholds" xml:"AlarmThresholds"`
}

func (s SetResellerUserAlarmThresholdRequest) String() string {
  return tea.Prettify(s)
}

func (s SetResellerUserAlarmThresholdRequest) GoString() string {
  return s.String()
}

func (s *SetResellerUserAlarmThresholdRequest) SetOwnerId(v int64) *SetResellerUserAlarmThresholdRequest {
  s.OwnerId = &v
  return s
}

func (s *SetResellerUserAlarmThresholdRequest) SetAlarmType(v string) *SetResellerUserAlarmThresholdRequest {
  s.AlarmType = &v
  return s
}

func (s *SetResellerUserAlarmThresholdRequest) SetAlarmThresholds(v string) *SetResellerUserAlarmThresholdRequest {
  s.AlarmThresholds = &v
  return s
}

type SetResellerUserAlarmThresholdResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Data *bool `json:"Data" xml:"Data" require:"true"`
}

func (s SetResellerUserAlarmThresholdResponse) String() string {
  return tea.Prettify(s)
}

func (s SetResellerUserAlarmThresholdResponse) GoString() string {
  return s.String()
}

func (s *SetResellerUserAlarmThresholdResponse) SetRequestId(v string) *SetResellerUserAlarmThresholdResponse {
  s.RequestId = &v
  return s
}

func (s *SetResellerUserAlarmThresholdResponse) SetCode(v string) *SetResellerUserAlarmThresholdResponse {
  s.Code = &v
  return s
}

func (s *SetResellerUserAlarmThresholdResponse) SetMessage(v string) *SetResellerUserAlarmThresholdResponse {
  s.Message = &v
  return s
}

func (s *SetResellerUserAlarmThresholdResponse) SetSuccess(v bool) *SetResellerUserAlarmThresholdResponse {
  s.Success = &v
  return s
}

func (s *SetResellerUserAlarmThresholdResponse) SetData(v bool) *SetResellerUserAlarmThresholdResponse {
  s.Data = &v
  return s
}

type QueryAccountTransactionsRequest struct {
  TransactionNumber *string `json:"TransactionNumber" xml:"TransactionNumber"`
  RecordID *string `json:"RecordID" xml:"RecordID"`
  TransactionChannelSN *string `json:"TransactionChannelSN" xml:"TransactionChannelSN"`
  CreateTimeStart *string `json:"CreateTimeStart" xml:"CreateTimeStart"`
  CreateTimeEnd *string `json:"CreateTimeEnd" xml:"CreateTimeEnd"`
  PageNum *int `json:"PageNum" xml:"PageNum"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
}

func (s QueryAccountTransactionsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTransactionsRequest) GoString() string {
  return s.String()
}

func (s *QueryAccountTransactionsRequest) SetTransactionNumber(v string) *QueryAccountTransactionsRequest {
  s.TransactionNumber = &v
  return s
}

func (s *QueryAccountTransactionsRequest) SetRecordID(v string) *QueryAccountTransactionsRequest {
  s.RecordID = &v
  return s
}

func (s *QueryAccountTransactionsRequest) SetTransactionChannelSN(v string) *QueryAccountTransactionsRequest {
  s.TransactionChannelSN = &v
  return s
}

func (s *QueryAccountTransactionsRequest) SetCreateTimeStart(v string) *QueryAccountTransactionsRequest {
  s.CreateTimeStart = &v
  return s
}

func (s *QueryAccountTransactionsRequest) SetCreateTimeEnd(v string) *QueryAccountTransactionsRequest {
  s.CreateTimeEnd = &v
  return s
}

func (s *QueryAccountTransactionsRequest) SetPageNum(v int) *QueryAccountTransactionsRequest {
  s.PageNum = &v
  return s
}

func (s *QueryAccountTransactionsRequest) SetPageSize(v int) *QueryAccountTransactionsRequest {
  s.PageSize = &v
  return s
}

type QueryAccountTransactionsResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryAccountTransactionsResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryAccountTransactionsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTransactionsResponse) GoString() string {
  return s.String()
}

func (s *QueryAccountTransactionsResponse) SetRequestId(v string) *QueryAccountTransactionsResponse {
  s.RequestId = &v
  return s
}

func (s *QueryAccountTransactionsResponse) SetSuccess(v bool) *QueryAccountTransactionsResponse {
  s.Success = &v
  return s
}

func (s *QueryAccountTransactionsResponse) SetCode(v string) *QueryAccountTransactionsResponse {
  s.Code = &v
  return s
}

func (s *QueryAccountTransactionsResponse) SetMessage(v string) *QueryAccountTransactionsResponse {
  s.Message = &v
  return s
}

func (s *QueryAccountTransactionsResponse) SetData(v *QueryAccountTransactionsResponseData) *QueryAccountTransactionsResponse {
  s.Data = v
  return s
}

type QueryAccountTransactionsResponseData struct {
  AccountName *string `json:"AccountName" xml:"AccountName" require:"true"`
  TotalCount *int `json:"TotalCount" xml:"TotalCount" require:"true"`
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  AccountTransactionsList *QueryAccountTransactionsResponseDataAccountTransactionsList `json:"AccountTransactionsList" xml:"AccountTransactionsList" require:"true" type:"Struct"`
}

func (s QueryAccountTransactionsResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTransactionsResponseData) GoString() string {
  return s.String()
}

func (s *QueryAccountTransactionsResponseData) SetAccountName(v string) *QueryAccountTransactionsResponseData {
  s.AccountName = &v
  return s
}

func (s *QueryAccountTransactionsResponseData) SetTotalCount(v int) *QueryAccountTransactionsResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryAccountTransactionsResponseData) SetPageNum(v int) *QueryAccountTransactionsResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryAccountTransactionsResponseData) SetPageSize(v int) *QueryAccountTransactionsResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryAccountTransactionsResponseData) SetAccountTransactionsList(v *QueryAccountTransactionsResponseDataAccountTransactionsList) *QueryAccountTransactionsResponseData {
  s.AccountTransactionsList = v
  return s
}

type QueryAccountTransactionsResponseDataAccountTransactionsList struct {
  AccountTransactionsList []*QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList `json:"AccountTransactionsList" xml:"AccountTransactionsList" require:"true" type:"Repeated"`
}

func (s QueryAccountTransactionsResponseDataAccountTransactionsList) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTransactionsResponseDataAccountTransactionsList) GoString() string {
  return s.String()
}

func (s *QueryAccountTransactionsResponseDataAccountTransactionsList) SetAccountTransactionsList(v []*QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) *QueryAccountTransactionsResponseDataAccountTransactionsList {
  s.AccountTransactionsList = v
  return s
}

type QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList struct     {
  TransactionNumber *string `json:"TransactionNumber" xml:"TransactionNumber" require:"true"`
  TransactionTime *string `json:"TransactionTime" xml:"TransactionTime" require:"true"`
  TransactionFlow *string `json:"TransactionFlow" xml:"TransactionFlow" require:"true"`
  TransactionType *string `json:"TransactionType" xml:"TransactionType" require:"true"`
  TransactionChannel *string `json:"TransactionChannel" xml:"TransactionChannel" require:"true"`
  TransactionChannelSN *string `json:"TransactionChannelSN" xml:"TransactionChannelSN" require:"true"`
  FundType *string `json:"FundType" xml:"FundType" require:"true"`
  RecordID *string `json:"RecordID" xml:"RecordID" require:"true"`
  Remarks *string `json:"Remarks" xml:"Remarks" require:"true"`
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  Amount *string `json:"Amount" xml:"Amount" require:"true"`
  Balance *string `json:"Balance" xml:"Balance" require:"true"`
  TransactionAccount *string `json:"TransactionAccount" xml:"TransactionAccount" require:"true"`
}

func (s QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) GoString() string {
  return s.String()
}

func (s *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) SetTransactionNumber(v string) *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList {
  s.TransactionNumber = &v
  return s
}

func (s *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) SetTransactionTime(v string) *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList {
  s.TransactionTime = &v
  return s
}

func (s *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) SetTransactionFlow(v string) *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList {
  s.TransactionFlow = &v
  return s
}

func (s *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) SetTransactionType(v string) *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList {
  s.TransactionType = &v
  return s
}

func (s *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) SetTransactionChannel(v string) *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList {
  s.TransactionChannel = &v
  return s
}

func (s *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) SetTransactionChannelSN(v string) *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList {
  s.TransactionChannelSN = &v
  return s
}

func (s *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) SetFundType(v string) *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList {
  s.FundType = &v
  return s
}

func (s *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) SetRecordID(v string) *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList {
  s.RecordID = &v
  return s
}

func (s *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) SetRemarks(v string) *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList {
  s.Remarks = &v
  return s
}

func (s *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) SetBillingCycle(v string) *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList {
  s.BillingCycle = &v
  return s
}

func (s *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) SetAmount(v string) *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList {
  s.Amount = &v
  return s
}

func (s *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) SetBalance(v string) *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList {
  s.Balance = &v
  return s
}

func (s *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList) SetTransactionAccount(v string) *QueryAccountTransactionsResponseDataAccountTransactionsListAccountTransactionsList {
  s.TransactionAccount = &v
  return s
}

type UnsubscribeBillToOSSRequest struct {
  SubscribeType *string `json:"SubscribeType" xml:"SubscribeType" require:"true"`
  MultAccountRelSubscribe *string `json:"MultAccountRelSubscribe" xml:"MultAccountRelSubscribe"`
}

func (s UnsubscribeBillToOSSRequest) String() string {
  return tea.Prettify(s)
}

func (s UnsubscribeBillToOSSRequest) GoString() string {
  return s.String()
}

func (s *UnsubscribeBillToOSSRequest) SetSubscribeType(v string) *UnsubscribeBillToOSSRequest {
  s.SubscribeType = &v
  return s
}

func (s *UnsubscribeBillToOSSRequest) SetMultAccountRelSubscribe(v string) *UnsubscribeBillToOSSRequest {
  s.MultAccountRelSubscribe = &v
  return s
}

type UnsubscribeBillToOSSResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
}

func (s UnsubscribeBillToOSSResponse) String() string {
  return tea.Prettify(s)
}

func (s UnsubscribeBillToOSSResponse) GoString() string {
  return s.String()
}

func (s *UnsubscribeBillToOSSResponse) SetRequestId(v string) *UnsubscribeBillToOSSResponse {
  s.RequestId = &v
  return s
}

func (s *UnsubscribeBillToOSSResponse) SetSuccess(v bool) *UnsubscribeBillToOSSResponse {
  s.Success = &v
  return s
}

func (s *UnsubscribeBillToOSSResponse) SetCode(v string) *UnsubscribeBillToOSSResponse {
  s.Code = &v
  return s
}

func (s *UnsubscribeBillToOSSResponse) SetMessage(v string) *UnsubscribeBillToOSSResponse {
  s.Message = &v
  return s
}

type SubscribeBillToOSSRequest struct {
  SubscribeBucket *string `json:"SubscribeBucket" xml:"SubscribeBucket" require:"true"`
  SubscribeType *string `json:"SubscribeType" xml:"SubscribeType"`
  MultAccountRelSubscribe *string `json:"MultAccountRelSubscribe" xml:"MultAccountRelSubscribe"`
  BucketOwnerId *int64 `json:"BucketOwnerId" xml:"BucketOwnerId"`
}

func (s SubscribeBillToOSSRequest) String() string {
  return tea.Prettify(s)
}

func (s SubscribeBillToOSSRequest) GoString() string {
  return s.String()
}

func (s *SubscribeBillToOSSRequest) SetSubscribeBucket(v string) *SubscribeBillToOSSRequest {
  s.SubscribeBucket = &v
  return s
}

func (s *SubscribeBillToOSSRequest) SetSubscribeType(v string) *SubscribeBillToOSSRequest {
  s.SubscribeType = &v
  return s
}

func (s *SubscribeBillToOSSRequest) SetMultAccountRelSubscribe(v string) *SubscribeBillToOSSRequest {
  s.MultAccountRelSubscribe = &v
  return s
}

func (s *SubscribeBillToOSSRequest) SetBucketOwnerId(v int64) *SubscribeBillToOSSRequest {
  s.BucketOwnerId = &v
  return s
}

type SubscribeBillToOSSResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
}

func (s SubscribeBillToOSSResponse) String() string {
  return tea.Prettify(s)
}

func (s SubscribeBillToOSSResponse) GoString() string {
  return s.String()
}

func (s *SubscribeBillToOSSResponse) SetRequestId(v string) *SubscribeBillToOSSResponse {
  s.RequestId = &v
  return s
}

func (s *SubscribeBillToOSSResponse) SetSuccess(v bool) *SubscribeBillToOSSResponse {
  s.Success = &v
  return s
}

func (s *SubscribeBillToOSSResponse) SetCode(v string) *SubscribeBillToOSSResponse {
  s.Code = &v
  return s
}

func (s *SubscribeBillToOSSResponse) SetMessage(v string) *SubscribeBillToOSSResponse {
  s.Message = &v
  return s
}

type QueryUserOmsDataRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  Table *string `json:"Table" xml:"Table" require:"true"`
  DataType *string `json:"DataType" xml:"DataType" require:"true"`
  StartTime *string `json:"StartTime" xml:"StartTime" require:"true"`
  EndTime *string `json:"EndTime" xml:"EndTime" require:"true"`
  Marker *string `json:"Marker" xml:"Marker"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
}

func (s QueryUserOmsDataRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryUserOmsDataRequest) GoString() string {
  return s.String()
}

func (s *QueryUserOmsDataRequest) SetOwnerId(v int64) *QueryUserOmsDataRequest {
  s.OwnerId = &v
  return s
}

func (s *QueryUserOmsDataRequest) SetTable(v string) *QueryUserOmsDataRequest {
  s.Table = &v
  return s
}

func (s *QueryUserOmsDataRequest) SetDataType(v string) *QueryUserOmsDataRequest {
  s.DataType = &v
  return s
}

func (s *QueryUserOmsDataRequest) SetStartTime(v string) *QueryUserOmsDataRequest {
  s.StartTime = &v
  return s
}

func (s *QueryUserOmsDataRequest) SetEndTime(v string) *QueryUserOmsDataRequest {
  s.EndTime = &v
  return s
}

func (s *QueryUserOmsDataRequest) SetMarker(v string) *QueryUserOmsDataRequest {
  s.Marker = &v
  return s
}

func (s *QueryUserOmsDataRequest) SetPageSize(v int) *QueryUserOmsDataRequest {
  s.PageSize = &v
  return s
}

type QueryUserOmsDataResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryUserOmsDataResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryUserOmsDataResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryUserOmsDataResponse) GoString() string {
  return s.String()
}

func (s *QueryUserOmsDataResponse) SetRequestId(v string) *QueryUserOmsDataResponse {
  s.RequestId = &v
  return s
}

func (s *QueryUserOmsDataResponse) SetSuccess(v bool) *QueryUserOmsDataResponse {
  s.Success = &v
  return s
}

func (s *QueryUserOmsDataResponse) SetCode(v string) *QueryUserOmsDataResponse {
  s.Code = &v
  return s
}

func (s *QueryUserOmsDataResponse) SetMessage(v string) *QueryUserOmsDataResponse {
  s.Message = &v
  return s
}

func (s *QueryUserOmsDataResponse) SetData(v *QueryUserOmsDataResponseData) *QueryUserOmsDataResponse {
  s.Data = v
  return s
}

type QueryUserOmsDataResponseData struct {
  Marker *string `json:"Marker" xml:"Marker" require:"true"`
  HostId *string `json:"HostId" xml:"HostId" require:"true"`
  OmsData []map[string]interface{} `json:"OmsData" xml:"OmsData" require:"true" type:"Repeated"`
}

func (s QueryUserOmsDataResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryUserOmsDataResponseData) GoString() string {
  return s.String()
}

func (s *QueryUserOmsDataResponseData) SetMarker(v string) *QueryUserOmsDataResponseData {
  s.Marker = &v
  return s
}

func (s *QueryUserOmsDataResponseData) SetHostId(v string) *QueryUserOmsDataResponseData {
  s.HostId = &v
  return s
}

func (s *QueryUserOmsDataResponseData) SetOmsData(v []map[string]interface{}) *QueryUserOmsDataResponseData {
  s.OmsData = v
  return s
}

type CancelOrderRequest struct {
  OrderId *string `json:"OrderId" xml:"OrderId" require:"true"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
}

func (s CancelOrderRequest) String() string {
  return tea.Prettify(s)
}

func (s CancelOrderRequest) GoString() string {
  return s.String()
}

func (s *CancelOrderRequest) SetOrderId(v string) *CancelOrderRequest {
  s.OrderId = &v
  return s
}

func (s *CancelOrderRequest) SetOwnerId(v int64) *CancelOrderRequest {
  s.OwnerId = &v
  return s
}

type CancelOrderResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *CancelOrderResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s CancelOrderResponse) String() string {
  return tea.Prettify(s)
}

func (s CancelOrderResponse) GoString() string {
  return s.String()
}

func (s *CancelOrderResponse) SetRequestId(v string) *CancelOrderResponse {
  s.RequestId = &v
  return s
}

func (s *CancelOrderResponse) SetSuccess(v bool) *CancelOrderResponse {
  s.Success = &v
  return s
}

func (s *CancelOrderResponse) SetCode(v string) *CancelOrderResponse {
  s.Code = &v
  return s
}

func (s *CancelOrderResponse) SetMessage(v string) *CancelOrderResponse {
  s.Message = &v
  return s
}

func (s *CancelOrderResponse) SetData(v *CancelOrderResponseData) *CancelOrderResponse {
  s.Data = v
  return s
}

type CancelOrderResponseData struct {
  HostId *string `json:"HostId" xml:"HostId" require:"true"`
}

func (s CancelOrderResponseData) String() string {
  return tea.Prettify(s)
}

func (s CancelOrderResponseData) GoString() string {
  return s.String()
}

func (s *CancelOrderResponseData) SetHostId(v string) *CancelOrderResponseData {
  s.HostId = &v
  return s
}

type ApplyInvoiceRequest struct {
  InvoiceAmount *int64 `json:"InvoiceAmount" xml:"InvoiceAmount" require:"true"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  CustomerId *int64 `json:"CustomerId" xml:"CustomerId" require:"true"`
  AddressId *int64 `json:"AddressId" xml:"AddressId" require:"true"`
  InvoicingType *int `json:"InvoicingType" xml:"InvoicingType"`
  ProcessWay *int `json:"ProcessWay" xml:"ProcessWay"`
  ApplyUserNick *string `json:"ApplyUserNick" xml:"ApplyUserNick" require:"true"`
  SelectedIds []*int64 `json:"SelectedIds" xml:"SelectedIds" require:"true" type:"Repeated"`
  InvoiceByAmount *bool `json:"InvoiceByAmount" xml:"InvoiceByAmount"`
}

func (s ApplyInvoiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ApplyInvoiceRequest) GoString() string {
  return s.String()
}

func (s *ApplyInvoiceRequest) SetInvoiceAmount(v int64) *ApplyInvoiceRequest {
  s.InvoiceAmount = &v
  return s
}

func (s *ApplyInvoiceRequest) SetOwnerId(v int64) *ApplyInvoiceRequest {
  s.OwnerId = &v
  return s
}

func (s *ApplyInvoiceRequest) SetCustomerId(v int64) *ApplyInvoiceRequest {
  s.CustomerId = &v
  return s
}

func (s *ApplyInvoiceRequest) SetAddressId(v int64) *ApplyInvoiceRequest {
  s.AddressId = &v
  return s
}

func (s *ApplyInvoiceRequest) SetInvoicingType(v int) *ApplyInvoiceRequest {
  s.InvoicingType = &v
  return s
}

func (s *ApplyInvoiceRequest) SetProcessWay(v int) *ApplyInvoiceRequest {
  s.ProcessWay = &v
  return s
}

func (s *ApplyInvoiceRequest) SetApplyUserNick(v string) *ApplyInvoiceRequest {
  s.ApplyUserNick = &v
  return s
}

func (s *ApplyInvoiceRequest) SetSelectedIds(v []*int64) *ApplyInvoiceRequest {
  s.SelectedIds = v
  return s
}

func (s *ApplyInvoiceRequest) SetInvoiceByAmount(v bool) *ApplyInvoiceRequest {
  s.InvoiceByAmount = &v
  return s
}

type ApplyInvoiceResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *ApplyInvoiceResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s ApplyInvoiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ApplyInvoiceResponse) GoString() string {
  return s.String()
}

func (s *ApplyInvoiceResponse) SetRequestId(v string) *ApplyInvoiceResponse {
  s.RequestId = &v
  return s
}

func (s *ApplyInvoiceResponse) SetSuccess(v bool) *ApplyInvoiceResponse {
  s.Success = &v
  return s
}

func (s *ApplyInvoiceResponse) SetCode(v string) *ApplyInvoiceResponse {
  s.Code = &v
  return s
}

func (s *ApplyInvoiceResponse) SetMessage(v string) *ApplyInvoiceResponse {
  s.Message = &v
  return s
}

func (s *ApplyInvoiceResponse) SetData(v *ApplyInvoiceResponseData) *ApplyInvoiceResponse {
  s.Data = v
  return s
}

type ApplyInvoiceResponseData struct {
  InvoiceApplyId *int64 `json:"InvoiceApplyId" xml:"InvoiceApplyId" require:"true"`
}

func (s ApplyInvoiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ApplyInvoiceResponseData) GoString() string {
  return s.String()
}

func (s *ApplyInvoiceResponseData) SetInvoiceApplyId(v int64) *ApplyInvoiceResponseData {
  s.InvoiceApplyId = &v
  return s
}

type QueryCustomerAddressListRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
}

func (s QueryCustomerAddressListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAddressListRequest) GoString() string {
  return s.String()
}

func (s *QueryCustomerAddressListRequest) SetOwnerId(v int64) *QueryCustomerAddressListRequest {
  s.OwnerId = &v
  return s
}

type QueryCustomerAddressListResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryCustomerAddressListResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryCustomerAddressListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAddressListResponse) GoString() string {
  return s.String()
}

func (s *QueryCustomerAddressListResponse) SetRequestId(v string) *QueryCustomerAddressListResponse {
  s.RequestId = &v
  return s
}

func (s *QueryCustomerAddressListResponse) SetSuccess(v bool) *QueryCustomerAddressListResponse {
  s.Success = &v
  return s
}

func (s *QueryCustomerAddressListResponse) SetCode(v string) *QueryCustomerAddressListResponse {
  s.Code = &v
  return s
}

func (s *QueryCustomerAddressListResponse) SetMessage(v string) *QueryCustomerAddressListResponse {
  s.Message = &v
  return s
}

func (s *QueryCustomerAddressListResponse) SetData(v *QueryCustomerAddressListResponseData) *QueryCustomerAddressListResponse {
  s.Data = v
  return s
}

type QueryCustomerAddressListResponseData struct {
  CustomerInvoiceAddressList *QueryCustomerAddressListResponseDataCustomerInvoiceAddressList `json:"CustomerInvoiceAddressList" xml:"CustomerInvoiceAddressList" require:"true" type:"Struct"`
}

func (s QueryCustomerAddressListResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAddressListResponseData) GoString() string {
  return s.String()
}

func (s *QueryCustomerAddressListResponseData) SetCustomerInvoiceAddressList(v *QueryCustomerAddressListResponseDataCustomerInvoiceAddressList) *QueryCustomerAddressListResponseData {
  s.CustomerInvoiceAddressList = v
  return s
}

type QueryCustomerAddressListResponseDataCustomerInvoiceAddressList struct {
  CustomerInvoiceAddress []*QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress `json:"CustomerInvoiceAddress" xml:"CustomerInvoiceAddress" require:"true" type:"Repeated"`
}

func (s QueryCustomerAddressListResponseDataCustomerInvoiceAddressList) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAddressListResponseDataCustomerInvoiceAddressList) GoString() string {
  return s.String()
}

func (s *QueryCustomerAddressListResponseDataCustomerInvoiceAddressList) SetCustomerInvoiceAddress(v []*QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) *QueryCustomerAddressListResponseDataCustomerInvoiceAddressList {
  s.CustomerInvoiceAddress = v
  return s
}

type QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress struct     {
  Id *int64 `json:"Id" xml:"Id" require:"true"`
  UserId *int64 `json:"UserId" xml:"UserId" require:"true"`
  UserNick *string `json:"UserNick" xml:"UserNick" require:"true"`
  Addressee *string `json:"Addressee" xml:"Addressee" require:"true"`
  Province *string `json:"Province" xml:"Province" require:"true"`
  City *string `json:"City" xml:"City" require:"true"`
  County *string `json:"County" xml:"County" require:"true"`
  Street *string `json:"Street" xml:"Street" require:"true"`
  PostalCode *string `json:"PostalCode" xml:"PostalCode" require:"true"`
  Phone *string `json:"Phone" xml:"Phone" require:"true"`
  BizType *string `json:"BizType" xml:"BizType" require:"true"`
  DeliveryAddress *string `json:"DeliveryAddress" xml:"DeliveryAddress" require:"true"`
}

func (s QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) GoString() string {
  return s.String()
}

func (s *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetId(v int64) *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress {
  s.Id = &v
  return s
}

func (s *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetUserId(v int64) *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress {
  s.UserId = &v
  return s
}

func (s *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetUserNick(v string) *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress {
  s.UserNick = &v
  return s
}

func (s *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetAddressee(v string) *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress {
  s.Addressee = &v
  return s
}

func (s *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetProvince(v string) *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress {
  s.Province = &v
  return s
}

func (s *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetCity(v string) *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress {
  s.City = &v
  return s
}

func (s *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetCounty(v string) *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress {
  s.County = &v
  return s
}

func (s *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetStreet(v string) *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress {
  s.Street = &v
  return s
}

func (s *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetPostalCode(v string) *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress {
  s.PostalCode = &v
  return s
}

func (s *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetPhone(v string) *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress {
  s.Phone = &v
  return s
}

func (s *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetBizType(v string) *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress {
  s.BizType = &v
  return s
}

func (s *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetDeliveryAddress(v string) *QueryCustomerAddressListResponseDataCustomerInvoiceAddressListCustomerInvoiceAddress {
  s.DeliveryAddress = &v
  return s
}

type QueryEvaluateListRequest struct {
  Type *int `json:"Type" xml:"Type"`
  OutBizId *string `json:"OutBizId" xml:"OutBizId"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  PageNum *int `json:"PageNum" xml:"PageNum"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
  StartAmount *int64 `json:"StartAmount" xml:"StartAmount"`
  EndAmount *int64 `json:"EndAmount" xml:"EndAmount"`
  StartBizTime *string `json:"StartBizTime" xml:"StartBizTime"`
  EndBizTime *string `json:"EndBizTime" xml:"EndBizTime"`
  SortType *int `json:"SortType" xml:"SortType"`
  StartSearchTime *string `json:"StartSearchTime" xml:"StartSearchTime"`
  EndSearchTime *string `json:"EndSearchTime" xml:"EndSearchTime"`
  BillCycle *string `json:"BillCycle" xml:"BillCycle"`
  BizTypeList []*string `json:"BizTypeList" xml:"BizTypeList" type:"Repeated"`
}

func (s QueryEvaluateListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEvaluateListRequest) GoString() string {
  return s.String()
}

func (s *QueryEvaluateListRequest) SetType(v int) *QueryEvaluateListRequest {
  s.Type = &v
  return s
}

func (s *QueryEvaluateListRequest) SetOutBizId(v string) *QueryEvaluateListRequest {
  s.OutBizId = &v
  return s
}

func (s *QueryEvaluateListRequest) SetOwnerId(v int64) *QueryEvaluateListRequest {
  s.OwnerId = &v
  return s
}

func (s *QueryEvaluateListRequest) SetPageNum(v int) *QueryEvaluateListRequest {
  s.PageNum = &v
  return s
}

func (s *QueryEvaluateListRequest) SetPageSize(v int) *QueryEvaluateListRequest {
  s.PageSize = &v
  return s
}

func (s *QueryEvaluateListRequest) SetStartAmount(v int64) *QueryEvaluateListRequest {
  s.StartAmount = &v
  return s
}

func (s *QueryEvaluateListRequest) SetEndAmount(v int64) *QueryEvaluateListRequest {
  s.EndAmount = &v
  return s
}

func (s *QueryEvaluateListRequest) SetStartBizTime(v string) *QueryEvaluateListRequest {
  s.StartBizTime = &v
  return s
}

func (s *QueryEvaluateListRequest) SetEndBizTime(v string) *QueryEvaluateListRequest {
  s.EndBizTime = &v
  return s
}

func (s *QueryEvaluateListRequest) SetSortType(v int) *QueryEvaluateListRequest {
  s.SortType = &v
  return s
}

func (s *QueryEvaluateListRequest) SetStartSearchTime(v string) *QueryEvaluateListRequest {
  s.StartSearchTime = &v
  return s
}

func (s *QueryEvaluateListRequest) SetEndSearchTime(v string) *QueryEvaluateListRequest {
  s.EndSearchTime = &v
  return s
}

func (s *QueryEvaluateListRequest) SetBillCycle(v string) *QueryEvaluateListRequest {
  s.BillCycle = &v
  return s
}

func (s *QueryEvaluateListRequest) SetBizTypeList(v []*string) *QueryEvaluateListRequest {
  s.BizTypeList = v
  return s
}

type QueryEvaluateListResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryEvaluateListResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryEvaluateListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEvaluateListResponse) GoString() string {
  return s.String()
}

func (s *QueryEvaluateListResponse) SetRequestId(v string) *QueryEvaluateListResponse {
  s.RequestId = &v
  return s
}

func (s *QueryEvaluateListResponse) SetSuccess(v bool) *QueryEvaluateListResponse {
  s.Success = &v
  return s
}

func (s *QueryEvaluateListResponse) SetCode(v string) *QueryEvaluateListResponse {
  s.Code = &v
  return s
}

func (s *QueryEvaluateListResponse) SetMessage(v string) *QueryEvaluateListResponse {
  s.Message = &v
  return s
}

func (s *QueryEvaluateListResponse) SetData(v *QueryEvaluateListResponseData) *QueryEvaluateListResponse {
  s.Data = v
  return s
}

type QueryEvaluateListResponseData struct {
  HostId *string `json:"HostId" xml:"HostId" require:"true"`
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  TotalCount *int `json:"TotalCount" xml:"TotalCount" require:"true"`
  TotalInvoiceAmount *int64 `json:"TotalInvoiceAmount" xml:"TotalInvoiceAmount" require:"true"`
  TotalUnAppliedInvoiceAmount *int64 `json:"TotalUnAppliedInvoiceAmount" xml:"TotalUnAppliedInvoiceAmount" require:"true"`
  EvaluateList *QueryEvaluateListResponseDataEvaluateList `json:"EvaluateList" xml:"EvaluateList" require:"true" type:"Struct"`
}

func (s QueryEvaluateListResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryEvaluateListResponseData) GoString() string {
  return s.String()
}

func (s *QueryEvaluateListResponseData) SetHostId(v string) *QueryEvaluateListResponseData {
  s.HostId = &v
  return s
}

func (s *QueryEvaluateListResponseData) SetPageNum(v int) *QueryEvaluateListResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryEvaluateListResponseData) SetPageSize(v int) *QueryEvaluateListResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryEvaluateListResponseData) SetTotalCount(v int) *QueryEvaluateListResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryEvaluateListResponseData) SetTotalInvoiceAmount(v int64) *QueryEvaluateListResponseData {
  s.TotalInvoiceAmount = &v
  return s
}

func (s *QueryEvaluateListResponseData) SetTotalUnAppliedInvoiceAmount(v int64) *QueryEvaluateListResponseData {
  s.TotalUnAppliedInvoiceAmount = &v
  return s
}

func (s *QueryEvaluateListResponseData) SetEvaluateList(v *QueryEvaluateListResponseDataEvaluateList) *QueryEvaluateListResponseData {
  s.EvaluateList = v
  return s
}

type QueryEvaluateListResponseDataEvaluateList struct {
  Evaluate []*QueryEvaluateListResponseDataEvaluateListEvaluate `json:"Evaluate" xml:"Evaluate" require:"true" type:"Repeated"`
}

func (s QueryEvaluateListResponseDataEvaluateList) String() string {
  return tea.Prettify(s)
}

func (s QueryEvaluateListResponseDataEvaluateList) GoString() string {
  return s.String()
}

func (s *QueryEvaluateListResponseDataEvaluateList) SetEvaluate(v []*QueryEvaluateListResponseDataEvaluateListEvaluate) *QueryEvaluateListResponseDataEvaluateList {
  s.Evaluate = v
  return s
}

type QueryEvaluateListResponseDataEvaluateListEvaluate struct     {
  Id *int64 `json:"Id" xml:"Id" require:"true"`
  GmtCreate *string `json:"GmtCreate" xml:"GmtCreate" require:"true"`
  GmtModified *string `json:"GmtModified" xml:"GmtModified" require:"true"`
  UserId *int64 `json:"UserId" xml:"UserId" require:"true"`
  UserNick *string `json:"UserNick" xml:"UserNick" require:"true"`
  OutBizId *string `json:"OutBizId" xml:"OutBizId" require:"true"`
  BillId *int64 `json:"BillId" xml:"BillId" require:"true"`
  ItemId *int64 `json:"ItemId" xml:"ItemId" require:"true"`
  BillCycle *string `json:"BillCycle" xml:"BillCycle" require:"true"`
  BizType *string `json:"BizType" xml:"BizType" require:"true"`
  OriginalAmount *int64 `json:"OriginalAmount" xml:"OriginalAmount" require:"true"`
  PresentAmount *int64 `json:"PresentAmount" xml:"PresentAmount" require:"true"`
  CanInvoiceAmount *int64 `json:"CanInvoiceAmount" xml:"CanInvoiceAmount" require:"true"`
  InvoicedAmount *int64 `json:"InvoicedAmount" xml:"InvoicedAmount" require:"true"`
  OffsetCostAmount *int64 `json:"OffsetCostAmount" xml:"OffsetCostAmount" require:"true"`
  OffsetAcceptAmount *int64 `json:"OffsetAcceptAmount" xml:"OffsetAcceptAmount" require:"true"`
  Status *int `json:"Status" xml:"Status" require:"true"`
  OpId *string `json:"OpId" xml:"OpId" require:"true"`
  Name *string `json:"Name" xml:"Name" require:"true"`
  BizTime *string `json:"BizTime" xml:"BizTime" require:"true"`
  Type *int `json:"Type" xml:"Type" require:"true"`
}

func (s QueryEvaluateListResponseDataEvaluateListEvaluate) String() string {
  return tea.Prettify(s)
}

func (s QueryEvaluateListResponseDataEvaluateListEvaluate) GoString() string {
  return s.String()
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetId(v int64) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.Id = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetGmtCreate(v string) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.GmtCreate = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetGmtModified(v string) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.GmtModified = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetUserId(v int64) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.UserId = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetUserNick(v string) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.UserNick = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetOutBizId(v string) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.OutBizId = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetBillId(v int64) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.BillId = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetItemId(v int64) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.ItemId = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetBillCycle(v string) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.BillCycle = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetBizType(v string) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.BizType = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetOriginalAmount(v int64) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.OriginalAmount = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetPresentAmount(v int64) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.PresentAmount = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetCanInvoiceAmount(v int64) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.CanInvoiceAmount = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetInvoicedAmount(v int64) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.InvoicedAmount = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetOffsetCostAmount(v int64) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.OffsetCostAmount = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetOffsetAcceptAmount(v int64) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.OffsetAcceptAmount = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetStatus(v int) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.Status = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetOpId(v string) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.OpId = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetName(v string) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.Name = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetBizTime(v string) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.BizTime = &v
  return s
}

func (s *QueryEvaluateListResponseDataEvaluateListEvaluate) SetType(v int) *QueryEvaluateListResponseDataEvaluateListEvaluate {
  s.Type = &v
  return s
}

type QueryInvoicingCustomerListRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
}

func (s QueryInvoicingCustomerListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryInvoicingCustomerListRequest) GoString() string {
  return s.String()
}

func (s *QueryInvoicingCustomerListRequest) SetOwnerId(v int64) *QueryInvoicingCustomerListRequest {
  s.OwnerId = &v
  return s
}

type QueryInvoicingCustomerListResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryInvoicingCustomerListResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryInvoicingCustomerListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryInvoicingCustomerListResponse) GoString() string {
  return s.String()
}

func (s *QueryInvoicingCustomerListResponse) SetRequestId(v string) *QueryInvoicingCustomerListResponse {
  s.RequestId = &v
  return s
}

func (s *QueryInvoicingCustomerListResponse) SetSuccess(v bool) *QueryInvoicingCustomerListResponse {
  s.Success = &v
  return s
}

func (s *QueryInvoicingCustomerListResponse) SetCode(v string) *QueryInvoicingCustomerListResponse {
  s.Code = &v
  return s
}

func (s *QueryInvoicingCustomerListResponse) SetMessage(v string) *QueryInvoicingCustomerListResponse {
  s.Message = &v
  return s
}

func (s *QueryInvoicingCustomerListResponse) SetData(v *QueryInvoicingCustomerListResponseData) *QueryInvoicingCustomerListResponse {
  s.Data = v
  return s
}

type QueryInvoicingCustomerListResponseData struct {
  CustomerInvoiceList *QueryInvoicingCustomerListResponseDataCustomerInvoiceList `json:"CustomerInvoiceList" xml:"CustomerInvoiceList" require:"true" type:"Struct"`
}

func (s QueryInvoicingCustomerListResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryInvoicingCustomerListResponseData) GoString() string {
  return s.String()
}

func (s *QueryInvoicingCustomerListResponseData) SetCustomerInvoiceList(v *QueryInvoicingCustomerListResponseDataCustomerInvoiceList) *QueryInvoicingCustomerListResponseData {
  s.CustomerInvoiceList = v
  return s
}

type QueryInvoicingCustomerListResponseDataCustomerInvoiceList struct {
  CustomerInvoice []*QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice `json:"CustomerInvoice" xml:"CustomerInvoice" require:"true" type:"Repeated"`
}

func (s QueryInvoicingCustomerListResponseDataCustomerInvoiceList) String() string {
  return tea.Prettify(s)
}

func (s QueryInvoicingCustomerListResponseDataCustomerInvoiceList) GoString() string {
  return s.String()
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceList) SetCustomerInvoice(v []*QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) *QueryInvoicingCustomerListResponseDataCustomerInvoiceList {
  s.CustomerInvoice = v
  return s
}

type QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice struct     {
  Id *int64 `json:"Id" xml:"Id" require:"true"`
  UserId *int64 `json:"UserId" xml:"UserId" require:"true"`
  UserNick *string `json:"UserNick" xml:"UserNick" require:"true"`
  InvoiceTitle *string `json:"InvoiceTitle" xml:"InvoiceTitle" require:"true"`
  CustomerType *int64 `json:"CustomerType" xml:"CustomerType" require:"true"`
  TaxpayerType *int64 `json:"TaxpayerType" xml:"TaxpayerType" require:"true"`
  Bank *string `json:"Bank" xml:"Bank" require:"true"`
  BankNo *string `json:"BankNo" xml:"BankNo" require:"true"`
  OperatingLicenseAddress *string `json:"OperatingLicenseAddress" xml:"OperatingLicenseAddress" require:"true"`
  OperatingLicensePhone *string `json:"OperatingLicensePhone" xml:"OperatingLicensePhone" require:"true"`
  RegisterNo *string `json:"RegisterNo" xml:"RegisterNo" require:"true"`
  StartCycle *int64 `json:"StartCycle" xml:"StartCycle" require:"true"`
  Status *int64 `json:"Status" xml:"Status" require:"true"`
  GmtCreate *string `json:"GmtCreate" xml:"GmtCreate" require:"true"`
  TaxationLicense *string `json:"TaxationLicense" xml:"TaxationLicense" require:"true"`
  AdjustType *int64 `json:"AdjustType" xml:"AdjustType" require:"true"`
  EndCycle *int64 `json:"EndCycle" xml:"EndCycle" require:"true"`
  TitleChangeInstructions *string `json:"TitleChangeInstructions" xml:"TitleChangeInstructions" require:"true"`
  IssueType *int64 `json:"IssueType" xml:"IssueType" require:"true"`
  Type *int64 `json:"Type" xml:"Type" require:"true"`
  DefaultRemark *string `json:"DefaultRemark" xml:"DefaultRemark" require:"true"`
}

func (s QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) String() string {
  return tea.Prettify(s)
}

func (s QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) GoString() string {
  return s.String()
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetId(v int64) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.Id = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetUserId(v int64) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.UserId = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetUserNick(v string) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.UserNick = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetInvoiceTitle(v string) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.InvoiceTitle = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetCustomerType(v int64) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.CustomerType = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetTaxpayerType(v int64) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.TaxpayerType = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetBank(v string) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.Bank = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetBankNo(v string) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.BankNo = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetOperatingLicenseAddress(v string) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.OperatingLicenseAddress = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetOperatingLicensePhone(v string) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.OperatingLicensePhone = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetRegisterNo(v string) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.RegisterNo = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetStartCycle(v int64) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.StartCycle = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetStatus(v int64) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.Status = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetGmtCreate(v string) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.GmtCreate = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetTaxationLicense(v string) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.TaxationLicense = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetAdjustType(v int64) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.AdjustType = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetEndCycle(v int64) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.EndCycle = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetTitleChangeInstructions(v string) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.TitleChangeInstructions = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetIssueType(v int64) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.IssueType = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetType(v int64) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.Type = &v
  return s
}

func (s *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice) SetDefaultRemark(v string) *QueryInvoicingCustomerListResponseDataCustomerInvoiceListCustomerInvoice {
  s.DefaultRemark = &v
  return s
}

type QueryBillOverviewRequest struct {
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType"`
}

func (s QueryBillOverviewRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBillOverviewRequest) GoString() string {
  return s.String()
}

func (s *QueryBillOverviewRequest) SetBillingCycle(v string) *QueryBillOverviewRequest {
  s.BillingCycle = &v
  return s
}

func (s *QueryBillOverviewRequest) SetProductCode(v string) *QueryBillOverviewRequest {
  s.ProductCode = &v
  return s
}

func (s *QueryBillOverviewRequest) SetProductType(v string) *QueryBillOverviewRequest {
  s.ProductType = &v
  return s
}

func (s *QueryBillOverviewRequest) SetSubscriptionType(v string) *QueryBillOverviewRequest {
  s.SubscriptionType = &v
  return s
}

type QueryBillOverviewResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryBillOverviewResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryBillOverviewResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBillOverviewResponse) GoString() string {
  return s.String()
}

func (s *QueryBillOverviewResponse) SetRequestId(v string) *QueryBillOverviewResponse {
  s.RequestId = &v
  return s
}

func (s *QueryBillOverviewResponse) SetSuccess(v bool) *QueryBillOverviewResponse {
  s.Success = &v
  return s
}

func (s *QueryBillOverviewResponse) SetCode(v string) *QueryBillOverviewResponse {
  s.Code = &v
  return s
}

func (s *QueryBillOverviewResponse) SetMessage(v string) *QueryBillOverviewResponse {
  s.Message = &v
  return s
}

func (s *QueryBillOverviewResponse) SetData(v *QueryBillOverviewResponseData) *QueryBillOverviewResponse {
  s.Data = v
  return s
}

type QueryBillOverviewResponseData struct {
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  AccountID *string `json:"AccountID" xml:"AccountID" require:"true"`
  AccountName *string `json:"AccountName" xml:"AccountName" require:"true"`
  Items *QueryBillOverviewResponseDataItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s QueryBillOverviewResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryBillOverviewResponseData) GoString() string {
  return s.String()
}

func (s *QueryBillOverviewResponseData) SetBillingCycle(v string) *QueryBillOverviewResponseData {
  s.BillingCycle = &v
  return s
}

func (s *QueryBillOverviewResponseData) SetAccountID(v string) *QueryBillOverviewResponseData {
  s.AccountID = &v
  return s
}

func (s *QueryBillOverviewResponseData) SetAccountName(v string) *QueryBillOverviewResponseData {
  s.AccountName = &v
  return s
}

func (s *QueryBillOverviewResponseData) SetItems(v *QueryBillOverviewResponseDataItems) *QueryBillOverviewResponseData {
  s.Items = v
  return s
}

type QueryBillOverviewResponseDataItems struct {
  Item []*QueryBillOverviewResponseDataItemsItem `json:"Item" xml:"Item" require:"true" type:"Repeated"`
}

func (s QueryBillOverviewResponseDataItems) String() string {
  return tea.Prettify(s)
}

func (s QueryBillOverviewResponseDataItems) GoString() string {
  return s.String()
}

func (s *QueryBillOverviewResponseDataItems) SetItem(v []*QueryBillOverviewResponseDataItemsItem) *QueryBillOverviewResponseDataItems {
  s.Item = v
  return s
}

type QueryBillOverviewResponseDataItemsItem struct     {
  Item *string `json:"Item" xml:"Item" require:"true"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType" require:"true"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  ProductName *string `json:"ProductName" xml:"ProductName" require:"true"`
  ProductDetail *string `json:"ProductDetail" xml:"ProductDetail" require:"true"`
  PretaxGrossAmount *float32 `json:"PretaxGrossAmount" xml:"PretaxGrossAmount" require:"true"`
  InvoiceDiscount *float32 `json:"InvoiceDiscount" xml:"InvoiceDiscount" require:"true"`
  DeductedByCoupons *float32 `json:"DeductedByCoupons" xml:"DeductedByCoupons" require:"true"`
  PretaxAmount *float32 `json:"PretaxAmount" xml:"PretaxAmount" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  PaymentAmount *float32 `json:"PaymentAmount" xml:"PaymentAmount" require:"true"`
  OutstandingAmount *float32 `json:"OutstandingAmount" xml:"OutstandingAmount" require:"true"`
  DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons" xml:"DeductedByCashCoupons" require:"true"`
  DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard" xml:"DeductedByPrepaidCard" require:"true"`
  PretaxAmountLocal *float32 `json:"PretaxAmountLocal" xml:"PretaxAmountLocal" require:"true"`
  Tax *float32 `json:"Tax" xml:"Tax" require:"true"`
  AfterTaxAmount *float32 `json:"AfterTaxAmount" xml:"AfterTaxAmount" require:"true"`
  PaymentCurrency *string `json:"PaymentCurrency" xml:"PaymentCurrency" require:"true"`
  RoundDownDiscount *string `json:"RoundDownDiscount" xml:"RoundDownDiscount" require:"true"`
}

func (s QueryBillOverviewResponseDataItemsItem) String() string {
  return tea.Prettify(s)
}

func (s QueryBillOverviewResponseDataItemsItem) GoString() string {
  return s.String()
}

func (s *QueryBillOverviewResponseDataItemsItem) SetItem(v string) *QueryBillOverviewResponseDataItemsItem {
  s.Item = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetProductCode(v string) *QueryBillOverviewResponseDataItemsItem {
  s.ProductCode = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetProductType(v string) *QueryBillOverviewResponseDataItemsItem {
  s.ProductType = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetSubscriptionType(v string) *QueryBillOverviewResponseDataItemsItem {
  s.SubscriptionType = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetProductName(v string) *QueryBillOverviewResponseDataItemsItem {
  s.ProductName = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetProductDetail(v string) *QueryBillOverviewResponseDataItemsItem {
  s.ProductDetail = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetPretaxGrossAmount(v float32) *QueryBillOverviewResponseDataItemsItem {
  s.PretaxGrossAmount = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetInvoiceDiscount(v float32) *QueryBillOverviewResponseDataItemsItem {
  s.InvoiceDiscount = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetDeductedByCoupons(v float32) *QueryBillOverviewResponseDataItemsItem {
  s.DeductedByCoupons = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetPretaxAmount(v float32) *QueryBillOverviewResponseDataItemsItem {
  s.PretaxAmount = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetCurrency(v string) *QueryBillOverviewResponseDataItemsItem {
  s.Currency = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetPaymentAmount(v float32) *QueryBillOverviewResponseDataItemsItem {
  s.PaymentAmount = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetOutstandingAmount(v float32) *QueryBillOverviewResponseDataItemsItem {
  s.OutstandingAmount = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetDeductedByCashCoupons(v float32) *QueryBillOverviewResponseDataItemsItem {
  s.DeductedByCashCoupons = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetDeductedByPrepaidCard(v float32) *QueryBillOverviewResponseDataItemsItem {
  s.DeductedByPrepaidCard = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetPretaxAmountLocal(v float32) *QueryBillOverviewResponseDataItemsItem {
  s.PretaxAmountLocal = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetTax(v float32) *QueryBillOverviewResponseDataItemsItem {
  s.Tax = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetAfterTaxAmount(v float32) *QueryBillOverviewResponseDataItemsItem {
  s.AfterTaxAmount = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetPaymentCurrency(v string) *QueryBillOverviewResponseDataItemsItem {
  s.PaymentCurrency = &v
  return s
}

func (s *QueryBillOverviewResponseDataItemsItem) SetRoundDownDiscount(v string) *QueryBillOverviewResponseDataItemsItem {
  s.RoundDownDiscount = &v
  return s
}

type QueryBillRequest struct {
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  Type *string `json:"Type" xml:"Type"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType"`
  IsHideZeroCharge *bool `json:"IsHideZeroCharge" xml:"IsHideZeroCharge"`
  IsDisplayLocalCurrency *bool `json:"IsDisplayLocalCurrency" xml:"IsDisplayLocalCurrency"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  PageNum *int `json:"PageNum" xml:"PageNum"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
}

func (s QueryBillRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBillRequest) GoString() string {
  return s.String()
}

func (s *QueryBillRequest) SetBillingCycle(v string) *QueryBillRequest {
  s.BillingCycle = &v
  return s
}

func (s *QueryBillRequest) SetType(v string) *QueryBillRequest {
  s.Type = &v
  return s
}

func (s *QueryBillRequest) SetProductCode(v string) *QueryBillRequest {
  s.ProductCode = &v
  return s
}

func (s *QueryBillRequest) SetProductType(v string) *QueryBillRequest {
  s.ProductType = &v
  return s
}

func (s *QueryBillRequest) SetSubscriptionType(v string) *QueryBillRequest {
  s.SubscriptionType = &v
  return s
}

func (s *QueryBillRequest) SetIsHideZeroCharge(v bool) *QueryBillRequest {
  s.IsHideZeroCharge = &v
  return s
}

func (s *QueryBillRequest) SetIsDisplayLocalCurrency(v bool) *QueryBillRequest {
  s.IsDisplayLocalCurrency = &v
  return s
}

func (s *QueryBillRequest) SetOwnerId(v int64) *QueryBillRequest {
  s.OwnerId = &v
  return s
}

func (s *QueryBillRequest) SetPageNum(v int) *QueryBillRequest {
  s.PageNum = &v
  return s
}

func (s *QueryBillRequest) SetPageSize(v int) *QueryBillRequest {
  s.PageSize = &v
  return s
}

type QueryBillResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryBillResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryBillResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBillResponse) GoString() string {
  return s.String()
}

func (s *QueryBillResponse) SetRequestId(v string) *QueryBillResponse {
  s.RequestId = &v
  return s
}

func (s *QueryBillResponse) SetSuccess(v bool) *QueryBillResponse {
  s.Success = &v
  return s
}

func (s *QueryBillResponse) SetCode(v string) *QueryBillResponse {
  s.Code = &v
  return s
}

func (s *QueryBillResponse) SetMessage(v string) *QueryBillResponse {
  s.Message = &v
  return s
}

func (s *QueryBillResponse) SetData(v *QueryBillResponseData) *QueryBillResponse {
  s.Data = v
  return s
}

type QueryBillResponseData struct {
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  AccountID *string `json:"AccountID" xml:"AccountID" require:"true"`
  AccountName *string `json:"AccountName" xml:"AccountName" require:"true"`
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  TotalCount *int `json:"TotalCount" xml:"TotalCount" require:"true"`
  Items *QueryBillResponseDataItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s QueryBillResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryBillResponseData) GoString() string {
  return s.String()
}

func (s *QueryBillResponseData) SetBillingCycle(v string) *QueryBillResponseData {
  s.BillingCycle = &v
  return s
}

func (s *QueryBillResponseData) SetAccountID(v string) *QueryBillResponseData {
  s.AccountID = &v
  return s
}

func (s *QueryBillResponseData) SetAccountName(v string) *QueryBillResponseData {
  s.AccountName = &v
  return s
}

func (s *QueryBillResponseData) SetPageNum(v int) *QueryBillResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryBillResponseData) SetPageSize(v int) *QueryBillResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryBillResponseData) SetTotalCount(v int) *QueryBillResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryBillResponseData) SetItems(v *QueryBillResponseDataItems) *QueryBillResponseData {
  s.Items = v
  return s
}

type QueryBillResponseDataItems struct {
  Item []*QueryBillResponseDataItemsItem `json:"Item" xml:"Item" require:"true" type:"Repeated"`
}

func (s QueryBillResponseDataItems) String() string {
  return tea.Prettify(s)
}

func (s QueryBillResponseDataItems) GoString() string {
  return s.String()
}

func (s *QueryBillResponseDataItems) SetItem(v []*QueryBillResponseDataItemsItem) *QueryBillResponseDataItems {
  s.Item = v
  return s
}

type QueryBillResponseDataItemsItem struct     {
  RecordID *string `json:"RecordID" xml:"RecordID" require:"true"`
  Item *string `json:"Item" xml:"Item" require:"true"`
  OwnerID *string `json:"OwnerID" xml:"OwnerID" require:"true"`
  UsageStartTime *string `json:"UsageStartTime" xml:"UsageStartTime" require:"true"`
  UsageEndTime *string `json:"UsageEndTime" xml:"UsageEndTime" require:"true"`
  PaymentTime *string `json:"PaymentTime" xml:"PaymentTime" require:"true"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType" require:"true"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  ProductName *string `json:"ProductName" xml:"ProductName" require:"true"`
  ProductDetail *string `json:"ProductDetail" xml:"ProductDetail" require:"true"`
  PretaxGrossAmount *float32 `json:"PretaxGrossAmount" xml:"PretaxGrossAmount" require:"true"`
  DeductedByCoupons *float32 `json:"DeductedByCoupons" xml:"DeductedByCoupons" require:"true"`
  InvoiceDiscount *float32 `json:"InvoiceDiscount" xml:"InvoiceDiscount" require:"true"`
  PretaxAmount *float32 `json:"PretaxAmount" xml:"PretaxAmount" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  PretaxAmountLocal *float32 `json:"PretaxAmountLocal" xml:"PretaxAmountLocal" require:"true"`
  Tax *float32 `json:"Tax" xml:"Tax" require:"true"`
  PaymentAmount *float32 `json:"PaymentAmount" xml:"PaymentAmount" require:"true"`
  DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons" xml:"DeductedByCashCoupons" require:"true"`
  DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard" xml:"DeductedByPrepaidCard" require:"true"`
  OutstandingAmount *float32 `json:"OutstandingAmount" xml:"OutstandingAmount" require:"true"`
  AfterTaxAmount *float32 `json:"AfterTaxAmount" xml:"AfterTaxAmount" require:"true"`
  Status *string `json:"Status" xml:"Status" require:"true"`
  PaymentCurrency *string `json:"PaymentCurrency" xml:"PaymentCurrency" require:"true"`
  PaymentTransactionID *string `json:"PaymentTransactionID" xml:"PaymentTransactionID" require:"true"`
  RoundDownDiscount *string `json:"RoundDownDiscount" xml:"RoundDownDiscount" require:"true"`
  SubOrderId *string `json:"SubOrderId" xml:"SubOrderId" require:"true"`
}

func (s QueryBillResponseDataItemsItem) String() string {
  return tea.Prettify(s)
}

func (s QueryBillResponseDataItemsItem) GoString() string {
  return s.String()
}

func (s *QueryBillResponseDataItemsItem) SetRecordID(v string) *QueryBillResponseDataItemsItem {
  s.RecordID = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetItem(v string) *QueryBillResponseDataItemsItem {
  s.Item = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetOwnerID(v string) *QueryBillResponseDataItemsItem {
  s.OwnerID = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetUsageStartTime(v string) *QueryBillResponseDataItemsItem {
  s.UsageStartTime = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetUsageEndTime(v string) *QueryBillResponseDataItemsItem {
  s.UsageEndTime = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetPaymentTime(v string) *QueryBillResponseDataItemsItem {
  s.PaymentTime = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetProductCode(v string) *QueryBillResponseDataItemsItem {
  s.ProductCode = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetProductType(v string) *QueryBillResponseDataItemsItem {
  s.ProductType = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetSubscriptionType(v string) *QueryBillResponseDataItemsItem {
  s.SubscriptionType = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetProductName(v string) *QueryBillResponseDataItemsItem {
  s.ProductName = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetProductDetail(v string) *QueryBillResponseDataItemsItem {
  s.ProductDetail = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetPretaxGrossAmount(v float32) *QueryBillResponseDataItemsItem {
  s.PretaxGrossAmount = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetDeductedByCoupons(v float32) *QueryBillResponseDataItemsItem {
  s.DeductedByCoupons = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetInvoiceDiscount(v float32) *QueryBillResponseDataItemsItem {
  s.InvoiceDiscount = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetPretaxAmount(v float32) *QueryBillResponseDataItemsItem {
  s.PretaxAmount = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetCurrency(v string) *QueryBillResponseDataItemsItem {
  s.Currency = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetPretaxAmountLocal(v float32) *QueryBillResponseDataItemsItem {
  s.PretaxAmountLocal = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetTax(v float32) *QueryBillResponseDataItemsItem {
  s.Tax = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetPaymentAmount(v float32) *QueryBillResponseDataItemsItem {
  s.PaymentAmount = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetDeductedByCashCoupons(v float32) *QueryBillResponseDataItemsItem {
  s.DeductedByCashCoupons = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetDeductedByPrepaidCard(v float32) *QueryBillResponseDataItemsItem {
  s.DeductedByPrepaidCard = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetOutstandingAmount(v float32) *QueryBillResponseDataItemsItem {
  s.OutstandingAmount = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetAfterTaxAmount(v float32) *QueryBillResponseDataItemsItem {
  s.AfterTaxAmount = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetStatus(v string) *QueryBillResponseDataItemsItem {
  s.Status = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetPaymentCurrency(v string) *QueryBillResponseDataItemsItem {
  s.PaymentCurrency = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetPaymentTransactionID(v string) *QueryBillResponseDataItemsItem {
  s.PaymentTransactionID = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetRoundDownDiscount(v string) *QueryBillResponseDataItemsItem {
  s.RoundDownDiscount = &v
  return s
}

func (s *QueryBillResponseDataItemsItem) SetSubOrderId(v string) *QueryBillResponseDataItemsItem {
  s.SubOrderId = &v
  return s
}

type QueryInstanceBillRequest struct {
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  IsBillingItem *bool `json:"IsBillingItem" xml:"IsBillingItem"`
  PageNum *int `json:"PageNum" xml:"PageNum"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
  IsHideZeroCharge *bool `json:"IsHideZeroCharge" xml:"IsHideZeroCharge"`
  BillingDate *string `json:"BillingDate" xml:"BillingDate"`
  Granularity *string `json:"Granularity" xml:"Granularity"`
}

func (s QueryInstanceBillRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceBillRequest) GoString() string {
  return s.String()
}

func (s *QueryInstanceBillRequest) SetBillingCycle(v string) *QueryInstanceBillRequest {
  s.BillingCycle = &v
  return s
}

func (s *QueryInstanceBillRequest) SetProductCode(v string) *QueryInstanceBillRequest {
  s.ProductCode = &v
  return s
}

func (s *QueryInstanceBillRequest) SetProductType(v string) *QueryInstanceBillRequest {
  s.ProductType = &v
  return s
}

func (s *QueryInstanceBillRequest) SetSubscriptionType(v string) *QueryInstanceBillRequest {
  s.SubscriptionType = &v
  return s
}

func (s *QueryInstanceBillRequest) SetOwnerId(v int64) *QueryInstanceBillRequest {
  s.OwnerId = &v
  return s
}

func (s *QueryInstanceBillRequest) SetIsBillingItem(v bool) *QueryInstanceBillRequest {
  s.IsBillingItem = &v
  return s
}

func (s *QueryInstanceBillRequest) SetPageNum(v int) *QueryInstanceBillRequest {
  s.PageNum = &v
  return s
}

func (s *QueryInstanceBillRequest) SetPageSize(v int) *QueryInstanceBillRequest {
  s.PageSize = &v
  return s
}

func (s *QueryInstanceBillRequest) SetIsHideZeroCharge(v bool) *QueryInstanceBillRequest {
  s.IsHideZeroCharge = &v
  return s
}

func (s *QueryInstanceBillRequest) SetBillingDate(v string) *QueryInstanceBillRequest {
  s.BillingDate = &v
  return s
}

func (s *QueryInstanceBillRequest) SetGranularity(v string) *QueryInstanceBillRequest {
  s.Granularity = &v
  return s
}

type QueryInstanceBillResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryInstanceBillResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryInstanceBillResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceBillResponse) GoString() string {
  return s.String()
}

func (s *QueryInstanceBillResponse) SetRequestId(v string) *QueryInstanceBillResponse {
  s.RequestId = &v
  return s
}

func (s *QueryInstanceBillResponse) SetSuccess(v bool) *QueryInstanceBillResponse {
  s.Success = &v
  return s
}

func (s *QueryInstanceBillResponse) SetCode(v string) *QueryInstanceBillResponse {
  s.Code = &v
  return s
}

func (s *QueryInstanceBillResponse) SetMessage(v string) *QueryInstanceBillResponse {
  s.Message = &v
  return s
}

func (s *QueryInstanceBillResponse) SetData(v *QueryInstanceBillResponseData) *QueryInstanceBillResponse {
  s.Data = v
  return s
}

type QueryInstanceBillResponseData struct {
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  AccountID *string `json:"AccountID" xml:"AccountID" require:"true"`
  AccountName *string `json:"AccountName" xml:"AccountName" require:"true"`
  TotalCount *int `json:"TotalCount" xml:"TotalCount" require:"true"`
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  Items *QueryInstanceBillResponseDataItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s QueryInstanceBillResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceBillResponseData) GoString() string {
  return s.String()
}

func (s *QueryInstanceBillResponseData) SetBillingCycle(v string) *QueryInstanceBillResponseData {
  s.BillingCycle = &v
  return s
}

func (s *QueryInstanceBillResponseData) SetAccountID(v string) *QueryInstanceBillResponseData {
  s.AccountID = &v
  return s
}

func (s *QueryInstanceBillResponseData) SetAccountName(v string) *QueryInstanceBillResponseData {
  s.AccountName = &v
  return s
}

func (s *QueryInstanceBillResponseData) SetTotalCount(v int) *QueryInstanceBillResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryInstanceBillResponseData) SetPageNum(v int) *QueryInstanceBillResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryInstanceBillResponseData) SetPageSize(v int) *QueryInstanceBillResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryInstanceBillResponseData) SetItems(v *QueryInstanceBillResponseDataItems) *QueryInstanceBillResponseData {
  s.Items = v
  return s
}

type QueryInstanceBillResponseDataItems struct {
  Item []*QueryInstanceBillResponseDataItemsItem `json:"Item" xml:"Item" require:"true" type:"Repeated"`
}

func (s QueryInstanceBillResponseDataItems) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceBillResponseDataItems) GoString() string {
  return s.String()
}

func (s *QueryInstanceBillResponseDataItems) SetItem(v []*QueryInstanceBillResponseDataItemsItem) *QueryInstanceBillResponseDataItems {
  s.Item = v
  return s
}

type QueryInstanceBillResponseDataItemsItem struct     {
  InstanceID *string `json:"InstanceID" xml:"InstanceID" require:"true"`
  BillingType *string `json:"BillingType" xml:"BillingType" require:"true"`
  CostUnit *string `json:"CostUnit" xml:"CostUnit" require:"true"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType" require:"true"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  ProductName *string `json:"ProductName" xml:"ProductName" require:"true"`
  ProductDetail *string `json:"ProductDetail" xml:"ProductDetail" require:"true"`
  OwnerID *string `json:"OwnerID" xml:"OwnerID" require:"true"`
  BillingItem *string `json:"BillingItem" xml:"BillingItem" require:"true"`
  ListPrice *string `json:"ListPrice" xml:"ListPrice" require:"true"`
  ListPriceUnit *string `json:"ListPriceUnit" xml:"ListPriceUnit" require:"true"`
  Usage *string `json:"Usage" xml:"Usage" require:"true"`
  UsageUnit *string `json:"UsageUnit" xml:"UsageUnit" require:"true"`
  DeductedByResourcePackage *string `json:"DeductedByResourcePackage" xml:"DeductedByResourcePackage" require:"true"`
  PretaxGrossAmount *float32 `json:"PretaxGrossAmount" xml:"PretaxGrossAmount" require:"true"`
  InvoiceDiscount *float32 `json:"InvoiceDiscount" xml:"InvoiceDiscount" require:"true"`
  DeductedByCoupons *float32 `json:"DeductedByCoupons" xml:"DeductedByCoupons" require:"true"`
  PretaxAmount *float32 `json:"PretaxAmount" xml:"PretaxAmount" require:"true"`
  DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons" xml:"DeductedByCashCoupons" require:"true"`
  DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard" xml:"DeductedByPrepaidCard" require:"true"`
  PaymentAmount *float32 `json:"PaymentAmount" xml:"PaymentAmount" require:"true"`
  OutstandingAmount *float32 `json:"OutstandingAmount" xml:"OutstandingAmount" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  NickName *string `json:"NickName" xml:"NickName" require:"true"`
  ResourceGroup *string `json:"ResourceGroup" xml:"ResourceGroup" require:"true"`
  Tag *string `json:"Tag" xml:"Tag" require:"true"`
  InstanceConfig *string `json:"InstanceConfig" xml:"InstanceConfig" require:"true"`
  InstanceSpec *string `json:"InstanceSpec" xml:"InstanceSpec" require:"true"`
  InternetIP *string `json:"InternetIP" xml:"InternetIP" require:"true"`
  IntranetIP *string `json:"IntranetIP" xml:"IntranetIP" require:"true"`
  Region *string `json:"Region" xml:"Region" require:"true"`
  Zone *string `json:"Zone" xml:"Zone" require:"true"`
  Item *string `json:"Item" xml:"Item" require:"true"`
  ServicePeriod *string `json:"ServicePeriod" xml:"ServicePeriod" require:"true"`
  BillingDate *string `json:"BillingDate" xml:"BillingDate" require:"true"`
}

func (s QueryInstanceBillResponseDataItemsItem) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceBillResponseDataItemsItem) GoString() string {
  return s.String()
}

func (s *QueryInstanceBillResponseDataItemsItem) SetInstanceID(v string) *QueryInstanceBillResponseDataItemsItem {
  s.InstanceID = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetBillingType(v string) *QueryInstanceBillResponseDataItemsItem {
  s.BillingType = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetCostUnit(v string) *QueryInstanceBillResponseDataItemsItem {
  s.CostUnit = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetProductCode(v string) *QueryInstanceBillResponseDataItemsItem {
  s.ProductCode = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetProductType(v string) *QueryInstanceBillResponseDataItemsItem {
  s.ProductType = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetSubscriptionType(v string) *QueryInstanceBillResponseDataItemsItem {
  s.SubscriptionType = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetProductName(v string) *QueryInstanceBillResponseDataItemsItem {
  s.ProductName = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetProductDetail(v string) *QueryInstanceBillResponseDataItemsItem {
  s.ProductDetail = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetOwnerID(v string) *QueryInstanceBillResponseDataItemsItem {
  s.OwnerID = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetBillingItem(v string) *QueryInstanceBillResponseDataItemsItem {
  s.BillingItem = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetListPrice(v string) *QueryInstanceBillResponseDataItemsItem {
  s.ListPrice = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetListPriceUnit(v string) *QueryInstanceBillResponseDataItemsItem {
  s.ListPriceUnit = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetUsage(v string) *QueryInstanceBillResponseDataItemsItem {
  s.Usage = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetUsageUnit(v string) *QueryInstanceBillResponseDataItemsItem {
  s.UsageUnit = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetDeductedByResourcePackage(v string) *QueryInstanceBillResponseDataItemsItem {
  s.DeductedByResourcePackage = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetPretaxGrossAmount(v float32) *QueryInstanceBillResponseDataItemsItem {
  s.PretaxGrossAmount = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetInvoiceDiscount(v float32) *QueryInstanceBillResponseDataItemsItem {
  s.InvoiceDiscount = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetDeductedByCoupons(v float32) *QueryInstanceBillResponseDataItemsItem {
  s.DeductedByCoupons = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetPretaxAmount(v float32) *QueryInstanceBillResponseDataItemsItem {
  s.PretaxAmount = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetDeductedByCashCoupons(v float32) *QueryInstanceBillResponseDataItemsItem {
  s.DeductedByCashCoupons = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetDeductedByPrepaidCard(v float32) *QueryInstanceBillResponseDataItemsItem {
  s.DeductedByPrepaidCard = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetPaymentAmount(v float32) *QueryInstanceBillResponseDataItemsItem {
  s.PaymentAmount = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetOutstandingAmount(v float32) *QueryInstanceBillResponseDataItemsItem {
  s.OutstandingAmount = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetCurrency(v string) *QueryInstanceBillResponseDataItemsItem {
  s.Currency = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetNickName(v string) *QueryInstanceBillResponseDataItemsItem {
  s.NickName = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetResourceGroup(v string) *QueryInstanceBillResponseDataItemsItem {
  s.ResourceGroup = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetTag(v string) *QueryInstanceBillResponseDataItemsItem {
  s.Tag = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetInstanceConfig(v string) *QueryInstanceBillResponseDataItemsItem {
  s.InstanceConfig = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetInstanceSpec(v string) *QueryInstanceBillResponseDataItemsItem {
  s.InstanceSpec = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetInternetIP(v string) *QueryInstanceBillResponseDataItemsItem {
  s.InternetIP = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetIntranetIP(v string) *QueryInstanceBillResponseDataItemsItem {
  s.IntranetIP = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetRegion(v string) *QueryInstanceBillResponseDataItemsItem {
  s.Region = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetZone(v string) *QueryInstanceBillResponseDataItemsItem {
  s.Zone = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetItem(v string) *QueryInstanceBillResponseDataItemsItem {
  s.Item = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetServicePeriod(v string) *QueryInstanceBillResponseDataItemsItem {
  s.ServicePeriod = &v
  return s
}

func (s *QueryInstanceBillResponseDataItemsItem) SetBillingDate(v string) *QueryInstanceBillResponseDataItemsItem {
  s.BillingDate = &v
  return s
}

type EnableBillGenerationRequest struct {
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId" require:"true"`
}

func (s EnableBillGenerationRequest) String() string {
  return tea.Prettify(s)
}

func (s EnableBillGenerationRequest) GoString() string {
  return s.String()
}

func (s *EnableBillGenerationRequest) SetProductCode(v string) *EnableBillGenerationRequest {
  s.ProductCode = &v
  return s
}

func (s *EnableBillGenerationRequest) SetOwnerId(v int64) *EnableBillGenerationRequest {
  s.OwnerId = &v
  return s
}

type EnableBillGenerationResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *EnableBillGenerationResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s EnableBillGenerationResponse) String() string {
  return tea.Prettify(s)
}

func (s EnableBillGenerationResponse) GoString() string {
  return s.String()
}

func (s *EnableBillGenerationResponse) SetRequestId(v string) *EnableBillGenerationResponse {
  s.RequestId = &v
  return s
}

func (s *EnableBillGenerationResponse) SetSuccess(v bool) *EnableBillGenerationResponse {
  s.Success = &v
  return s
}

func (s *EnableBillGenerationResponse) SetCode(v string) *EnableBillGenerationResponse {
  s.Code = &v
  return s
}

func (s *EnableBillGenerationResponse) SetMessage(v string) *EnableBillGenerationResponse {
  s.Message = &v
  return s
}

func (s *EnableBillGenerationResponse) SetData(v *EnableBillGenerationResponseData) *EnableBillGenerationResponse {
  s.Data = v
  return s
}

type EnableBillGenerationResponseData struct {
  Boolean *bool `json:"Boolean" xml:"Boolean" require:"true"`
}

func (s EnableBillGenerationResponseData) String() string {
  return tea.Prettify(s)
}

func (s EnableBillGenerationResponseData) GoString() string {
  return s.String()
}

func (s *EnableBillGenerationResponseData) SetBoolean(v bool) *EnableBillGenerationResponseData {
  s.Boolean = &v
  return s
}

type QueryRedeemRequest struct {
  ExpiryTimeStart *string `json:"ExpiryTimeStart" xml:"ExpiryTimeStart"`
  ExpiryTimeEnd *string `json:"ExpiryTimeEnd" xml:"ExpiryTimeEnd"`
  EffectiveOrNot *bool `json:"EffectiveOrNot" xml:"EffectiveOrNot"`
  PageNum *int `json:"PageNum" xml:"PageNum"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
}

func (s QueryRedeemRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRedeemRequest) GoString() string {
  return s.String()
}

func (s *QueryRedeemRequest) SetExpiryTimeStart(v string) *QueryRedeemRequest {
  s.ExpiryTimeStart = &v
  return s
}

func (s *QueryRedeemRequest) SetExpiryTimeEnd(v string) *QueryRedeemRequest {
  s.ExpiryTimeEnd = &v
  return s
}

func (s *QueryRedeemRequest) SetEffectiveOrNot(v bool) *QueryRedeemRequest {
  s.EffectiveOrNot = &v
  return s
}

func (s *QueryRedeemRequest) SetPageNum(v int) *QueryRedeemRequest {
  s.PageNum = &v
  return s
}

func (s *QueryRedeemRequest) SetPageSize(v int) *QueryRedeemRequest {
  s.PageSize = &v
  return s
}

type QueryRedeemResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryRedeemResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryRedeemResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRedeemResponse) GoString() string {
  return s.String()
}

func (s *QueryRedeemResponse) SetRequestId(v string) *QueryRedeemResponse {
  s.RequestId = &v
  return s
}

func (s *QueryRedeemResponse) SetSuccess(v bool) *QueryRedeemResponse {
  s.Success = &v
  return s
}

func (s *QueryRedeemResponse) SetCode(v string) *QueryRedeemResponse {
  s.Code = &v
  return s
}

func (s *QueryRedeemResponse) SetMessage(v string) *QueryRedeemResponse {
  s.Message = &v
  return s
}

func (s *QueryRedeemResponse) SetData(v *QueryRedeemResponseData) *QueryRedeemResponse {
  s.Data = v
  return s
}

type QueryRedeemResponseData struct {
  PageNum *int64 `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int64 `json:"PageSize" xml:"PageSize" require:"true"`
  TotalCount *int64 `json:"TotalCount" xml:"TotalCount" require:"true"`
  Redeem *QueryRedeemResponseDataRedeem `json:"Redeem" xml:"Redeem" require:"true" type:"Struct"`
}

func (s QueryRedeemResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryRedeemResponseData) GoString() string {
  return s.String()
}

func (s *QueryRedeemResponseData) SetPageNum(v int64) *QueryRedeemResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryRedeemResponseData) SetPageSize(v int64) *QueryRedeemResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryRedeemResponseData) SetTotalCount(v int64) *QueryRedeemResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryRedeemResponseData) SetRedeem(v *QueryRedeemResponseDataRedeem) *QueryRedeemResponseData {
  s.Redeem = v
  return s
}

type QueryRedeemResponseDataRedeem struct {
  Redeem []*QueryRedeemResponseDataRedeemRedeem `json:"Redeem" xml:"Redeem" require:"true" type:"Repeated"`
}

func (s QueryRedeemResponseDataRedeem) String() string {
  return tea.Prettify(s)
}

func (s QueryRedeemResponseDataRedeem) GoString() string {
  return s.String()
}

func (s *QueryRedeemResponseDataRedeem) SetRedeem(v []*QueryRedeemResponseDataRedeemRedeem) *QueryRedeemResponseDataRedeem {
  s.Redeem = v
  return s
}

type QueryRedeemResponseDataRedeemRedeem struct     {
  RedeemId *string `json:"RedeemId" xml:"RedeemId" require:"true"`
  RedeemNo *string `json:"RedeemNo" xml:"RedeemNo" require:"true"`
  Status *string `json:"Status" xml:"Status" require:"true"`
  GrantedTime *string `json:"GrantedTime" xml:"GrantedTime" require:"true"`
  EffectiveTime *string `json:"EffectiveTime" xml:"EffectiveTime" require:"true"`
  ExpiryTime *string `json:"ExpiryTime" xml:"ExpiryTime" require:"true"`
  NominalValue *string `json:"NominalValue" xml:"NominalValue" require:"true"`
  Balance *string `json:"Balance" xml:"Balance" require:"true"`
  ApplicableProducts *string `json:"ApplicableProducts" xml:"ApplicableProducts" require:"true"`
  Specification *string `json:"Specification" xml:"Specification" require:"true"`
}

func (s QueryRedeemResponseDataRedeemRedeem) String() string {
  return tea.Prettify(s)
}

func (s QueryRedeemResponseDataRedeemRedeem) GoString() string {
  return s.String()
}

func (s *QueryRedeemResponseDataRedeemRedeem) SetRedeemId(v string) *QueryRedeemResponseDataRedeemRedeem {
  s.RedeemId = &v
  return s
}

func (s *QueryRedeemResponseDataRedeemRedeem) SetRedeemNo(v string) *QueryRedeemResponseDataRedeemRedeem {
  s.RedeemNo = &v
  return s
}

func (s *QueryRedeemResponseDataRedeemRedeem) SetStatus(v string) *QueryRedeemResponseDataRedeemRedeem {
  s.Status = &v
  return s
}

func (s *QueryRedeemResponseDataRedeemRedeem) SetGrantedTime(v string) *QueryRedeemResponseDataRedeemRedeem {
  s.GrantedTime = &v
  return s
}

func (s *QueryRedeemResponseDataRedeemRedeem) SetEffectiveTime(v string) *QueryRedeemResponseDataRedeemRedeem {
  s.EffectiveTime = &v
  return s
}

func (s *QueryRedeemResponseDataRedeemRedeem) SetExpiryTime(v string) *QueryRedeemResponseDataRedeemRedeem {
  s.ExpiryTime = &v
  return s
}

func (s *QueryRedeemResponseDataRedeemRedeem) SetNominalValue(v string) *QueryRedeemResponseDataRedeemRedeem {
  s.NominalValue = &v
  return s
}

func (s *QueryRedeemResponseDataRedeemRedeem) SetBalance(v string) *QueryRedeemResponseDataRedeemRedeem {
  s.Balance = &v
  return s
}

func (s *QueryRedeemResponseDataRedeemRedeem) SetApplicableProducts(v string) *QueryRedeemResponseDataRedeemRedeem {
  s.ApplicableProducts = &v
  return s
}

func (s *QueryRedeemResponseDataRedeemRedeem) SetSpecification(v string) *QueryRedeemResponseDataRedeemRedeem {
  s.Specification = &v
  return s
}

type ConvertChargeTypeRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  Period *int `json:"Period" xml:"Period"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  InstanceId *string `json:"InstanceId" xml:"InstanceId" require:"true"`
}

func (s ConvertChargeTypeRequest) String() string {
  return tea.Prettify(s)
}

func (s ConvertChargeTypeRequest) GoString() string {
  return s.String()
}

func (s *ConvertChargeTypeRequest) SetOwnerId(v int64) *ConvertChargeTypeRequest {
  s.OwnerId = &v
  return s
}

func (s *ConvertChargeTypeRequest) SetProductType(v string) *ConvertChargeTypeRequest {
  s.ProductType = &v
  return s
}

func (s *ConvertChargeTypeRequest) SetSubscriptionType(v string) *ConvertChargeTypeRequest {
  s.SubscriptionType = &v
  return s
}

func (s *ConvertChargeTypeRequest) SetPeriod(v int) *ConvertChargeTypeRequest {
  s.Period = &v
  return s
}

func (s *ConvertChargeTypeRequest) SetProductCode(v string) *ConvertChargeTypeRequest {
  s.ProductCode = &v
  return s
}

func (s *ConvertChargeTypeRequest) SetInstanceId(v string) *ConvertChargeTypeRequest {
  s.InstanceId = &v
  return s
}

type ConvertChargeTypeResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *ConvertChargeTypeResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s ConvertChargeTypeResponse) String() string {
  return tea.Prettify(s)
}

func (s ConvertChargeTypeResponse) GoString() string {
  return s.String()
}

func (s *ConvertChargeTypeResponse) SetRequestId(v string) *ConvertChargeTypeResponse {
  s.RequestId = &v
  return s
}

func (s *ConvertChargeTypeResponse) SetSuccess(v bool) *ConvertChargeTypeResponse {
  s.Success = &v
  return s
}

func (s *ConvertChargeTypeResponse) SetCode(v string) *ConvertChargeTypeResponse {
  s.Code = &v
  return s
}

func (s *ConvertChargeTypeResponse) SetMessage(v string) *ConvertChargeTypeResponse {
  s.Message = &v
  return s
}

func (s *ConvertChargeTypeResponse) SetData(v *ConvertChargeTypeResponseData) *ConvertChargeTypeResponse {
  s.Data = v
  return s
}

type ConvertChargeTypeResponseData struct {
  OrderId *string `json:"OrderId" xml:"OrderId" require:"true"`
}

func (s ConvertChargeTypeResponseData) String() string {
  return tea.Prettify(s)
}

func (s ConvertChargeTypeResponseData) GoString() string {
  return s.String()
}

func (s *ConvertChargeTypeResponseData) SetOrderId(v string) *ConvertChargeTypeResponseData {
  s.OrderId = &v
  return s
}

type CreateInstanceRequest struct {
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  Parameter []*CreateInstanceRequestParameter `json:"Parameter" xml:"Parameter" type:"Repeated"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  Period *int `json:"Period" xml:"Period"`
  RenewalStatus *string `json:"RenewalStatus" xml:"RenewalStatus"`
  RenewPeriod *int `json:"RenewPeriod" xml:"RenewPeriod"`
  ClientToken *string `json:"ClientToken" xml:"ClientToken"`
}

func (s CreateInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
  return s.String()
}

func (s *CreateInstanceRequest) SetProductCode(v string) *CreateInstanceRequest {
  s.ProductCode = &v
  return s
}

func (s *CreateInstanceRequest) SetParameter(v []*CreateInstanceRequestParameter) *CreateInstanceRequest {
  s.Parameter = v
  return s
}

func (s *CreateInstanceRequest) SetOwnerId(v int64) *CreateInstanceRequest {
  s.OwnerId = &v
  return s
}

func (s *CreateInstanceRequest) SetProductType(v string) *CreateInstanceRequest {
  s.ProductType = &v
  return s
}

func (s *CreateInstanceRequest) SetSubscriptionType(v string) *CreateInstanceRequest {
  s.SubscriptionType = &v
  return s
}

func (s *CreateInstanceRequest) SetPeriod(v int) *CreateInstanceRequest {
  s.Period = &v
  return s
}

func (s *CreateInstanceRequest) SetRenewalStatus(v string) *CreateInstanceRequest {
  s.RenewalStatus = &v
  return s
}

func (s *CreateInstanceRequest) SetRenewPeriod(v int) *CreateInstanceRequest {
  s.RenewPeriod = &v
  return s
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
  s.ClientToken = &v
  return s
}

type CreateInstanceRequestParameter struct     {
  Code *string `json:"Code" xml:"Code" require:"true"`
  Value *string `json:"Value" xml:"Value" require:"true"`
}

func (s CreateInstanceRequestParameter) String() string {
  return tea.Prettify(s)
}

func (s CreateInstanceRequestParameter) GoString() string {
  return s.String()
}

func (s *CreateInstanceRequestParameter) SetCode(v string) *CreateInstanceRequestParameter {
  s.Code = &v
  return s
}

func (s *CreateInstanceRequestParameter) SetValue(v string) *CreateInstanceRequestParameter {
  s.Value = &v
  return s
}

type CreateInstanceResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *CreateInstanceResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s CreateInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
  return s.String()
}

func (s *CreateInstanceResponse) SetRequestId(v string) *CreateInstanceResponse {
  s.RequestId = &v
  return s
}

func (s *CreateInstanceResponse) SetSuccess(v bool) *CreateInstanceResponse {
  s.Success = &v
  return s
}

func (s *CreateInstanceResponse) SetCode(v string) *CreateInstanceResponse {
  s.Code = &v
  return s
}

func (s *CreateInstanceResponse) SetMessage(v string) *CreateInstanceResponse {
  s.Message = &v
  return s
}

func (s *CreateInstanceResponse) SetData(v *CreateInstanceResponseData) *CreateInstanceResponse {
  s.Data = v
  return s
}

type CreateInstanceResponseData struct {
  InstanceId *string `json:"InstanceId" xml:"InstanceId" require:"true"`
  OrderId *string `json:"OrderId" xml:"OrderId" require:"true"`
}

func (s CreateInstanceResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateInstanceResponseData) GoString() string {
  return s.String()
}

func (s *CreateInstanceResponseData) SetInstanceId(v string) *CreateInstanceResponseData {
  s.InstanceId = &v
  return s
}

func (s *CreateInstanceResponseData) SetOrderId(v string) *CreateInstanceResponseData {
  s.OrderId = &v
  return s
}

type ModifyInstanceRequest struct {
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  ModifyType *string `json:"ModifyType" xml:"ModifyType" require:"true"`
  InstanceId *string `json:"InstanceId" xml:"InstanceId"`
  Parameter []*ModifyInstanceRequestParameter `json:"Parameter" xml:"Parameter" type:"Repeated"`
  ClientToken *string `json:"ClientToken" xml:"ClientToken"`
}

func (s ModifyInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s ModifyInstanceRequest) GoString() string {
  return s.String()
}

func (s *ModifyInstanceRequest) SetProductCode(v string) *ModifyInstanceRequest {
  s.ProductCode = &v
  return s
}

func (s *ModifyInstanceRequest) SetOwnerId(v int64) *ModifyInstanceRequest {
  s.OwnerId = &v
  return s
}

func (s *ModifyInstanceRequest) SetProductType(v string) *ModifyInstanceRequest {
  s.ProductType = &v
  return s
}

func (s *ModifyInstanceRequest) SetSubscriptionType(v string) *ModifyInstanceRequest {
  s.SubscriptionType = &v
  return s
}

func (s *ModifyInstanceRequest) SetModifyType(v string) *ModifyInstanceRequest {
  s.ModifyType = &v
  return s
}

func (s *ModifyInstanceRequest) SetInstanceId(v string) *ModifyInstanceRequest {
  s.InstanceId = &v
  return s
}

func (s *ModifyInstanceRequest) SetParameter(v []*ModifyInstanceRequestParameter) *ModifyInstanceRequest {
  s.Parameter = v
  return s
}

func (s *ModifyInstanceRequest) SetClientToken(v string) *ModifyInstanceRequest {
  s.ClientToken = &v
  return s
}

type ModifyInstanceRequestParameter struct     {
  Code *string `json:"Code" xml:"Code" require:"true"`
  Value *string `json:"Value" xml:"Value" require:"true"`
}

func (s ModifyInstanceRequestParameter) String() string {
  return tea.Prettify(s)
}

func (s ModifyInstanceRequestParameter) GoString() string {
  return s.String()
}

func (s *ModifyInstanceRequestParameter) SetCode(v string) *ModifyInstanceRequestParameter {
  s.Code = &v
  return s
}

func (s *ModifyInstanceRequestParameter) SetValue(v string) *ModifyInstanceRequestParameter {
  s.Value = &v
  return s
}

type ModifyInstanceResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *ModifyInstanceResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s ModifyInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s ModifyInstanceResponse) GoString() string {
  return s.String()
}

func (s *ModifyInstanceResponse) SetRequestId(v string) *ModifyInstanceResponse {
  s.RequestId = &v
  return s
}

func (s *ModifyInstanceResponse) SetSuccess(v bool) *ModifyInstanceResponse {
  s.Success = &v
  return s
}

func (s *ModifyInstanceResponse) SetCode(v string) *ModifyInstanceResponse {
  s.Code = &v
  return s
}

func (s *ModifyInstanceResponse) SetMessage(v string) *ModifyInstanceResponse {
  s.Message = &v
  return s
}

func (s *ModifyInstanceResponse) SetData(v *ModifyInstanceResponseData) *ModifyInstanceResponse {
  s.Data = v
  return s
}

type ModifyInstanceResponseData struct {
  HostId *string `json:"HostId" xml:"HostId" require:"true"`
  OrderId *string `json:"OrderId" xml:"OrderId" require:"true"`
}

func (s ModifyInstanceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ModifyInstanceResponseData) GoString() string {
  return s.String()
}

func (s *ModifyInstanceResponseData) SetHostId(v string) *ModifyInstanceResponseData {
  s.HostId = &v
  return s
}

func (s *ModifyInstanceResponseData) SetOrderId(v string) *ModifyInstanceResponseData {
  s.OrderId = &v
  return s
}

type DescribePricingModuleRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
}

func (s DescribePricingModuleRequest) String() string {
  return tea.Prettify(s)
}

func (s DescribePricingModuleRequest) GoString() string {
  return s.String()
}

func (s *DescribePricingModuleRequest) SetOwnerId(v int64) *DescribePricingModuleRequest {
  s.OwnerId = &v
  return s
}

func (s *DescribePricingModuleRequest) SetProductCode(v string) *DescribePricingModuleRequest {
  s.ProductCode = &v
  return s
}

func (s *DescribePricingModuleRequest) SetProductType(v string) *DescribePricingModuleRequest {
  s.ProductType = &v
  return s
}

func (s *DescribePricingModuleRequest) SetSubscriptionType(v string) *DescribePricingModuleRequest {
  s.SubscriptionType = &v
  return s
}

type DescribePricingModuleResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *DescribePricingModuleResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s DescribePricingModuleResponse) String() string {
  return tea.Prettify(s)
}

func (s DescribePricingModuleResponse) GoString() string {
  return s.String()
}

func (s *DescribePricingModuleResponse) SetRequestId(v string) *DescribePricingModuleResponse {
  s.RequestId = &v
  return s
}

func (s *DescribePricingModuleResponse) SetSuccess(v bool) *DescribePricingModuleResponse {
  s.Success = &v
  return s
}

func (s *DescribePricingModuleResponse) SetCode(v string) *DescribePricingModuleResponse {
  s.Code = &v
  return s
}

func (s *DescribePricingModuleResponse) SetMessage(v string) *DescribePricingModuleResponse {
  s.Message = &v
  return s
}

func (s *DescribePricingModuleResponse) SetData(v *DescribePricingModuleResponseData) *DescribePricingModuleResponse {
  s.Data = v
  return s
}

type DescribePricingModuleResponseData struct {
  ModuleList *DescribePricingModuleResponseDataModuleList `json:"ModuleList" xml:"ModuleList" require:"true" type:"Struct"`
  AttributeList *DescribePricingModuleResponseDataAttributeList `json:"AttributeList" xml:"AttributeList" require:"true" type:"Struct"`
}

func (s DescribePricingModuleResponseData) String() string {
  return tea.Prettify(s)
}

func (s DescribePricingModuleResponseData) GoString() string {
  return s.String()
}

func (s *DescribePricingModuleResponseData) SetModuleList(v *DescribePricingModuleResponseDataModuleList) *DescribePricingModuleResponseData {
  s.ModuleList = v
  return s
}

func (s *DescribePricingModuleResponseData) SetAttributeList(v *DescribePricingModuleResponseDataAttributeList) *DescribePricingModuleResponseData {
  s.AttributeList = v
  return s
}

type DescribePricingModuleResponseDataModuleList struct {
  Module []*DescribePricingModuleResponseDataModuleListModule `json:"Module" xml:"Module" require:"true" type:"Repeated"`
}

func (s DescribePricingModuleResponseDataModuleList) String() string {
  return tea.Prettify(s)
}

func (s DescribePricingModuleResponseDataModuleList) GoString() string {
  return s.String()
}

func (s *DescribePricingModuleResponseDataModuleList) SetModule(v []*DescribePricingModuleResponseDataModuleListModule) *DescribePricingModuleResponseDataModuleList {
  s.Module = v
  return s
}

type DescribePricingModuleResponseDataModuleListModule struct     {
  ModuleCode *string `json:"ModuleCode" xml:"ModuleCode" require:"true"`
  ModuleName *string `json:"ModuleName" xml:"ModuleName" require:"true"`
  PriceType *string `json:"PriceType" xml:"PriceType" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  ConfigList *DescribePricingModuleResponseDataModuleListModuleConfigList `json:"ConfigList" xml:"ConfigList" require:"true" type:"Struct"`
}

func (s DescribePricingModuleResponseDataModuleListModule) String() string {
  return tea.Prettify(s)
}

func (s DescribePricingModuleResponseDataModuleListModule) GoString() string {
  return s.String()
}

func (s *DescribePricingModuleResponseDataModuleListModule) SetModuleCode(v string) *DescribePricingModuleResponseDataModuleListModule {
  s.ModuleCode = &v
  return s
}

func (s *DescribePricingModuleResponseDataModuleListModule) SetModuleName(v string) *DescribePricingModuleResponseDataModuleListModule {
  s.ModuleName = &v
  return s
}

func (s *DescribePricingModuleResponseDataModuleListModule) SetPriceType(v string) *DescribePricingModuleResponseDataModuleListModule {
  s.PriceType = &v
  return s
}

func (s *DescribePricingModuleResponseDataModuleListModule) SetCurrency(v string) *DescribePricingModuleResponseDataModuleListModule {
  s.Currency = &v
  return s
}

func (s *DescribePricingModuleResponseDataModuleListModule) SetConfigList(v *DescribePricingModuleResponseDataModuleListModuleConfigList) *DescribePricingModuleResponseDataModuleListModule {
  s.ConfigList = v
  return s
}

type DescribePricingModuleResponseDataModuleListModuleConfigList struct {
  ConfigList []*string `json:"ConfigList" xml:"ConfigList" require:"true" type:"Repeated"`
}

func (s DescribePricingModuleResponseDataModuleListModuleConfigList) String() string {
  return tea.Prettify(s)
}

func (s DescribePricingModuleResponseDataModuleListModuleConfigList) GoString() string {
  return s.String()
}

func (s *DescribePricingModuleResponseDataModuleListModuleConfigList) SetConfigList(v []*string) *DescribePricingModuleResponseDataModuleListModuleConfigList {
  s.ConfigList = v
  return s
}

type DescribePricingModuleResponseDataAttributeList struct {
  Attribute []*DescribePricingModuleResponseDataAttributeListAttribute `json:"Attribute" xml:"Attribute" require:"true" type:"Repeated"`
}

func (s DescribePricingModuleResponseDataAttributeList) String() string {
  return tea.Prettify(s)
}

func (s DescribePricingModuleResponseDataAttributeList) GoString() string {
  return s.String()
}

func (s *DescribePricingModuleResponseDataAttributeList) SetAttribute(v []*DescribePricingModuleResponseDataAttributeListAttribute) *DescribePricingModuleResponseDataAttributeList {
  s.Attribute = v
  return s
}

type DescribePricingModuleResponseDataAttributeListAttribute struct     {
  Code *string `json:"Code" xml:"Code" require:"true"`
  Name *string `json:"Name" xml:"Name" require:"true"`
  Unit *string `json:"Unit" xml:"Unit" require:"true"`
  Values *DescribePricingModuleResponseDataAttributeListAttributeValues `json:"Values" xml:"Values" require:"true" type:"Struct"`
}

func (s DescribePricingModuleResponseDataAttributeListAttribute) String() string {
  return tea.Prettify(s)
}

func (s DescribePricingModuleResponseDataAttributeListAttribute) GoString() string {
  return s.String()
}

func (s *DescribePricingModuleResponseDataAttributeListAttribute) SetCode(v string) *DescribePricingModuleResponseDataAttributeListAttribute {
  s.Code = &v
  return s
}

func (s *DescribePricingModuleResponseDataAttributeListAttribute) SetName(v string) *DescribePricingModuleResponseDataAttributeListAttribute {
  s.Name = &v
  return s
}

func (s *DescribePricingModuleResponseDataAttributeListAttribute) SetUnit(v string) *DescribePricingModuleResponseDataAttributeListAttribute {
  s.Unit = &v
  return s
}

func (s *DescribePricingModuleResponseDataAttributeListAttribute) SetValues(v *DescribePricingModuleResponseDataAttributeListAttributeValues) *DescribePricingModuleResponseDataAttributeListAttribute {
  s.Values = v
  return s
}

type DescribePricingModuleResponseDataAttributeListAttributeValues struct {
  AttributeValue []*DescribePricingModuleResponseDataAttributeListAttributeValuesAttributeValue `json:"AttributeValue" xml:"AttributeValue" require:"true" type:"Repeated"`
}

func (s DescribePricingModuleResponseDataAttributeListAttributeValues) String() string {
  return tea.Prettify(s)
}

func (s DescribePricingModuleResponseDataAttributeListAttributeValues) GoString() string {
  return s.String()
}

func (s *DescribePricingModuleResponseDataAttributeListAttributeValues) SetAttributeValue(v []*DescribePricingModuleResponseDataAttributeListAttributeValuesAttributeValue) *DescribePricingModuleResponseDataAttributeListAttributeValues {
  s.AttributeValue = v
  return s
}

type DescribePricingModuleResponseDataAttributeListAttributeValuesAttributeValue struct     {
  Type *string `json:"Type" xml:"Type" require:"true"`
  Name *string `json:"Name" xml:"Name" require:"true"`
  Value *string `json:"Value" xml:"Value" require:"true"`
  Remark *string `json:"Remark" xml:"Remark" require:"true"`
}

func (s DescribePricingModuleResponseDataAttributeListAttributeValuesAttributeValue) String() string {
  return tea.Prettify(s)
}

func (s DescribePricingModuleResponseDataAttributeListAttributeValuesAttributeValue) GoString() string {
  return s.String()
}

func (s *DescribePricingModuleResponseDataAttributeListAttributeValuesAttributeValue) SetType(v string) *DescribePricingModuleResponseDataAttributeListAttributeValuesAttributeValue {
  s.Type = &v
  return s
}

func (s *DescribePricingModuleResponseDataAttributeListAttributeValuesAttributeValue) SetName(v string) *DescribePricingModuleResponseDataAttributeListAttributeValuesAttributeValue {
  s.Name = &v
  return s
}

func (s *DescribePricingModuleResponseDataAttributeListAttributeValuesAttributeValue) SetValue(v string) *DescribePricingModuleResponseDataAttributeListAttributeValuesAttributeValue {
  s.Value = &v
  return s
}

func (s *DescribePricingModuleResponseDataAttributeListAttributeValuesAttributeValue) SetRemark(v string) *DescribePricingModuleResponseDataAttributeListAttributeValuesAttributeValue {
  s.Remark = &v
  return s
}

type QueryProductListRequest struct {
  QueryTotalCount *bool `json:"QueryTotalCount" xml:"QueryTotalCount"`
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
}

func (s QueryProductListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryProductListRequest) GoString() string {
  return s.String()
}

func (s *QueryProductListRequest) SetQueryTotalCount(v bool) *QueryProductListRequest {
  s.QueryTotalCount = &v
  return s
}

func (s *QueryProductListRequest) SetPageNum(v int) *QueryProductListRequest {
  s.PageNum = &v
  return s
}

func (s *QueryProductListRequest) SetPageSize(v int) *QueryProductListRequest {
  s.PageSize = &v
  return s
}

type QueryProductListResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryProductListResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryProductListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryProductListResponse) GoString() string {
  return s.String()
}

func (s *QueryProductListResponse) SetRequestId(v string) *QueryProductListResponse {
  s.RequestId = &v
  return s
}

func (s *QueryProductListResponse) SetSuccess(v bool) *QueryProductListResponse {
  s.Success = &v
  return s
}

func (s *QueryProductListResponse) SetCode(v string) *QueryProductListResponse {
  s.Code = &v
  return s
}

func (s *QueryProductListResponse) SetMessage(v string) *QueryProductListResponse {
  s.Message = &v
  return s
}

func (s *QueryProductListResponse) SetData(v *QueryProductListResponseData) *QueryProductListResponse {
  s.Data = v
  return s
}

type QueryProductListResponseData struct {
  TotalCount *int `json:"TotalCount" xml:"TotalCount" require:"true"`
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  ProductList *QueryProductListResponseDataProductList `json:"ProductList" xml:"ProductList" require:"true" type:"Struct"`
}

func (s QueryProductListResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryProductListResponseData) GoString() string {
  return s.String()
}

func (s *QueryProductListResponseData) SetTotalCount(v int) *QueryProductListResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryProductListResponseData) SetPageNum(v int) *QueryProductListResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryProductListResponseData) SetPageSize(v int) *QueryProductListResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryProductListResponseData) SetProductList(v *QueryProductListResponseDataProductList) *QueryProductListResponseData {
  s.ProductList = v
  return s
}

type QueryProductListResponseDataProductList struct {
  Product []*QueryProductListResponseDataProductListProduct `json:"Product" xml:"Product" require:"true" type:"Repeated"`
}

func (s QueryProductListResponseDataProductList) String() string {
  return tea.Prettify(s)
}

func (s QueryProductListResponseDataProductList) GoString() string {
  return s.String()
}

func (s *QueryProductListResponseDataProductList) SetProduct(v []*QueryProductListResponseDataProductListProduct) *QueryProductListResponseDataProductList {
  s.Product = v
  return s
}

type QueryProductListResponseDataProductListProduct struct     {
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductName *string `json:"ProductName" xml:"ProductName" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType" require:"true"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
}

func (s QueryProductListResponseDataProductListProduct) String() string {
  return tea.Prettify(s)
}

func (s QueryProductListResponseDataProductListProduct) GoString() string {
  return s.String()
}

func (s *QueryProductListResponseDataProductListProduct) SetProductCode(v string) *QueryProductListResponseDataProductListProduct {
  s.ProductCode = &v
  return s
}

func (s *QueryProductListResponseDataProductListProduct) SetProductName(v string) *QueryProductListResponseDataProductListProduct {
  s.ProductName = &v
  return s
}

func (s *QueryProductListResponseDataProductListProduct) SetProductType(v string) *QueryProductListResponseDataProductListProduct {
  s.ProductType = &v
  return s
}

func (s *QueryProductListResponseDataProductListProduct) SetSubscriptionType(v string) *QueryProductListResponseDataProductListProduct {
  s.SubscriptionType = &v
  return s
}

type QueryInstanceGaapCostRequest struct {
  PageNum *int `json:"PageNum" xml:"PageNum"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType"`
}

func (s QueryInstanceGaapCostRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceGaapCostRequest) GoString() string {
  return s.String()
}

func (s *QueryInstanceGaapCostRequest) SetPageNum(v int) *QueryInstanceGaapCostRequest {
  s.PageNum = &v
  return s
}

func (s *QueryInstanceGaapCostRequest) SetPageSize(v int) *QueryInstanceGaapCostRequest {
  s.PageSize = &v
  return s
}

func (s *QueryInstanceGaapCostRequest) SetBillingCycle(v string) *QueryInstanceGaapCostRequest {
  s.BillingCycle = &v
  return s
}

func (s *QueryInstanceGaapCostRequest) SetProductCode(v string) *QueryInstanceGaapCostRequest {
  s.ProductCode = &v
  return s
}

func (s *QueryInstanceGaapCostRequest) SetProductType(v string) *QueryInstanceGaapCostRequest {
  s.ProductType = &v
  return s
}

func (s *QueryInstanceGaapCostRequest) SetSubscriptionType(v string) *QueryInstanceGaapCostRequest {
  s.SubscriptionType = &v
  return s
}

type QueryInstanceGaapCostResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryInstanceGaapCostResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryInstanceGaapCostResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceGaapCostResponse) GoString() string {
  return s.String()
}

func (s *QueryInstanceGaapCostResponse) SetRequestId(v string) *QueryInstanceGaapCostResponse {
  s.RequestId = &v
  return s
}

func (s *QueryInstanceGaapCostResponse) SetSuccess(v bool) *QueryInstanceGaapCostResponse {
  s.Success = &v
  return s
}

func (s *QueryInstanceGaapCostResponse) SetCode(v string) *QueryInstanceGaapCostResponse {
  s.Code = &v
  return s
}

func (s *QueryInstanceGaapCostResponse) SetMessage(v string) *QueryInstanceGaapCostResponse {
  s.Message = &v
  return s
}

func (s *QueryInstanceGaapCostResponse) SetData(v *QueryInstanceGaapCostResponseData) *QueryInstanceGaapCostResponse {
  s.Data = v
  return s
}

type QueryInstanceGaapCostResponseData struct {
  HostId *string `json:"HostId" xml:"HostId" require:"true"`
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  TotalCount *int `json:"TotalCount" xml:"TotalCount" require:"true"`
  Modules *QueryInstanceGaapCostResponseDataModules `json:"Modules" xml:"Modules" require:"true" type:"Struct"`
}

func (s QueryInstanceGaapCostResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceGaapCostResponseData) GoString() string {
  return s.String()
}

func (s *QueryInstanceGaapCostResponseData) SetHostId(v string) *QueryInstanceGaapCostResponseData {
  s.HostId = &v
  return s
}

func (s *QueryInstanceGaapCostResponseData) SetPageNum(v int) *QueryInstanceGaapCostResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryInstanceGaapCostResponseData) SetPageSize(v int) *QueryInstanceGaapCostResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryInstanceGaapCostResponseData) SetTotalCount(v int) *QueryInstanceGaapCostResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseData) SetModules(v *QueryInstanceGaapCostResponseDataModules) *QueryInstanceGaapCostResponseData {
  s.Modules = v
  return s
}

type QueryInstanceGaapCostResponseDataModules struct {
  Module []*QueryInstanceGaapCostResponseDataModulesModule `json:"Module" xml:"Module" require:"true" type:"Repeated"`
}

func (s QueryInstanceGaapCostResponseDataModules) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceGaapCostResponseDataModules) GoString() string {
  return s.String()
}

func (s *QueryInstanceGaapCostResponseDataModules) SetModule(v []*QueryInstanceGaapCostResponseDataModulesModule) *QueryInstanceGaapCostResponseDataModules {
  s.Module = v
  return s
}

type QueryInstanceGaapCostResponseDataModulesModule struct     {
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  InstanceID *string `json:"InstanceID" xml:"InstanceID" require:"true"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType" require:"true"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  Tag *string `json:"Tag" xml:"Tag" require:"true"`
  ResourceGroup *string `json:"ResourceGroup" xml:"ResourceGroup" require:"true"`
  AccountingUnit *string `json:"AccountingUnit" xml:"AccountingUnit" require:"true"`
  PayerAccount *string `json:"PayerAccount" xml:"PayerAccount" require:"true"`
  OwnerID *string `json:"OwnerID" xml:"OwnerID" require:"true"`
  Region *string `json:"Region" xml:"Region" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  PaymentCurrency *string `json:"PaymentCurrency" xml:"PaymentCurrency" require:"true"`
  OrderType *string `json:"OrderType" xml:"OrderType" require:"true"`
  PayTime *string `json:"PayTime" xml:"PayTime" require:"true"`
  PretaxGrossAmount *string `json:"PretaxGrossAmount" xml:"PretaxGrossAmount" require:"true"`
  PricingDiscount *string `json:"PricingDiscount" xml:"PricingDiscount" require:"true"`
  DeductedByCoupons *string `json:"DeductedByCoupons" xml:"DeductedByCoupons" require:"true"`
  PretaxAmount *string `json:"PretaxAmount" xml:"PretaxAmount" require:"true"`
  PretaxAmountLocal *string `json:"PretaxAmountLocal" xml:"PretaxAmountLocal" require:"true"`
  DeductedByCashCoupons *string `json:"DeductedByCashCoupons" xml:"DeductedByCashCoupons" require:"true"`
  DeductedByPrepaidCard *string `json:"DeductedByPrepaidCard" xml:"DeductedByPrepaidCard" require:"true"`
  PaymentAmount *string `json:"PaymentAmount" xml:"PaymentAmount" require:"true"`
  GaapPretaxGrossAmount *string `json:"GaapPretaxGrossAmount" xml:"GaapPretaxGrossAmount" require:"true"`
  GaapPricingDiscount *string `json:"GaapPricingDiscount" xml:"GaapPricingDiscount" require:"true"`
  GaapDeductedByCoupons *string `json:"GaapDeductedByCoupons" xml:"GaapDeductedByCoupons" require:"true"`
  GaapPretaxAmount *string `json:"GaapPretaxAmount" xml:"GaapPretaxAmount" require:"true"`
  GaapPretaxAmountLocal *string `json:"GaapPretaxAmountLocal" xml:"GaapPretaxAmountLocal" require:"true"`
  GaapDeductedByCashCoupons *string `json:"GaapDeductedByCashCoupons" xml:"GaapDeductedByCashCoupons" require:"true"`
  GaapDeductedByPrepaidCard *string `json:"GaapDeductedByPrepaidCard" xml:"GaapDeductedByPrepaidCard" require:"true"`
  GaapPaymentAmount *string `json:"GaapPaymentAmount" xml:"GaapPaymentAmount" require:"true"`
  MonthGaapPretaxGrossAmount *string `json:"MonthGaapPretaxGrossAmount" xml:"MonthGaapPretaxGrossAmount" require:"true"`
  MonthGaapPricingDiscount *string `json:"MonthGaapPricingDiscount" xml:"MonthGaapPricingDiscount" require:"true"`
  MonthGaapDeductedByCoupons *string `json:"MonthGaapDeductedByCoupons" xml:"MonthGaapDeductedByCoupons" require:"true"`
  MonthGaapPretaxAmount *string `json:"MonthGaapPretaxAmount" xml:"MonthGaapPretaxAmount" require:"true"`
  MonthGaapPretaxAmountLocal *string `json:"MonthGaapPretaxAmountLocal" xml:"MonthGaapPretaxAmountLocal" require:"true"`
  MonthGaapDeductedByCashCoupons *string `json:"MonthGaapDeductedByCashCoupons" xml:"MonthGaapDeductedByCashCoupons" require:"true"`
  MonthGaapDeductedByPrepaidCard *string `json:"MonthGaapDeductedByPrepaidCard" xml:"MonthGaapDeductedByPrepaidCard" require:"true"`
  MonthGaapPaymentAmount *string `json:"MonthGaapPaymentAmount" xml:"MonthGaapPaymentAmount" require:"true"`
  UnallocatedPaymentAmount *string `json:"UnallocatedPaymentAmount" xml:"UnallocatedPaymentAmount" require:"true"`
  UsageStartDate *string `json:"UsageStartDate" xml:"UsageStartDate" require:"true"`
  UsageEndDate *string `json:"UsageEndDate" xml:"UsageEndDate" require:"true"`
  BillType *string `json:"BillType" xml:"BillType" require:"true"`
  OrderId *string `json:"OrderId" xml:"OrderId" require:"true"`
  SubOrderId *string `json:"SubOrderId" xml:"SubOrderId" require:"true"`
  UnallocatedPretaxGrossAmount *string `json:"UnallocatedPretaxGrossAmount" xml:"UnallocatedPretaxGrossAmount" require:"true"`
  UnallocatedPricingDiscount *string `json:"UnallocatedPricingDiscount" xml:"UnallocatedPricingDiscount" require:"true"`
  UnallocatedDeductedByCoupons *string `json:"UnallocatedDeductedByCoupons" xml:"UnallocatedDeductedByCoupons" require:"true"`
  UnallocatedPretaxAmount *string `json:"UnallocatedPretaxAmount" xml:"UnallocatedPretaxAmount" require:"true"`
  UnallocatedPretaxAmountLocal *string `json:"UnallocatedPretaxAmountLocal" xml:"UnallocatedPretaxAmountLocal" require:"true"`
  UnallocatedDeductedByCashCoupons *string `json:"UnallocatedDeductedByCashCoupons" xml:"UnallocatedDeductedByCashCoupons" require:"true"`
  UnallocatedDeductedByPrepaidCard *string `json:"UnallocatedDeductedByPrepaidCard" xml:"UnallocatedDeductedByPrepaidCard" require:"true"`
}

func (s QueryInstanceGaapCostResponseDataModulesModule) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceGaapCostResponseDataModulesModule) GoString() string {
  return s.String()
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetBillingCycle(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.BillingCycle = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetInstanceID(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.InstanceID = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetProductCode(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.ProductCode = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetProductType(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.ProductType = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetSubscriptionType(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.SubscriptionType = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetTag(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.Tag = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetResourceGroup(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.ResourceGroup = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetAccountingUnit(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.AccountingUnit = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetPayerAccount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.PayerAccount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetOwnerID(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.OwnerID = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetRegion(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.Region = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetCurrency(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.Currency = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetPaymentCurrency(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.PaymentCurrency = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetOrderType(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.OrderType = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetPayTime(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.PayTime = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetPretaxGrossAmount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.PretaxGrossAmount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetPricingDiscount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.PricingDiscount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetDeductedByCoupons(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.DeductedByCoupons = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetPretaxAmount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.PretaxAmount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetPretaxAmountLocal(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.PretaxAmountLocal = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetDeductedByCashCoupons(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.DeductedByCashCoupons = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetDeductedByPrepaidCard(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.DeductedByPrepaidCard = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetPaymentAmount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.PaymentAmount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetGaapPretaxGrossAmount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.GaapPretaxGrossAmount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetGaapPricingDiscount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.GaapPricingDiscount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetGaapDeductedByCoupons(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.GaapDeductedByCoupons = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetGaapPretaxAmount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.GaapPretaxAmount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetGaapPretaxAmountLocal(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.GaapPretaxAmountLocal = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetGaapDeductedByCashCoupons(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.GaapDeductedByCashCoupons = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetGaapDeductedByPrepaidCard(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.GaapDeductedByPrepaidCard = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetGaapPaymentAmount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.GaapPaymentAmount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetMonthGaapPretaxGrossAmount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.MonthGaapPretaxGrossAmount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetMonthGaapPricingDiscount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.MonthGaapPricingDiscount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetMonthGaapDeductedByCoupons(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.MonthGaapDeductedByCoupons = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetMonthGaapPretaxAmount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.MonthGaapPretaxAmount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetMonthGaapPretaxAmountLocal(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.MonthGaapPretaxAmountLocal = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetMonthGaapDeductedByCashCoupons(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.MonthGaapDeductedByCashCoupons = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetMonthGaapDeductedByPrepaidCard(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.MonthGaapDeductedByPrepaidCard = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetMonthGaapPaymentAmount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.MonthGaapPaymentAmount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetUnallocatedPaymentAmount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.UnallocatedPaymentAmount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetUsageStartDate(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.UsageStartDate = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetUsageEndDate(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.UsageEndDate = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetBillType(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.BillType = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetOrderId(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.OrderId = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetSubOrderId(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.SubOrderId = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetUnallocatedPretaxGrossAmount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.UnallocatedPretaxGrossAmount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetUnallocatedPricingDiscount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.UnallocatedPricingDiscount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetUnallocatedDeductedByCoupons(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.UnallocatedDeductedByCoupons = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetUnallocatedPretaxAmount(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.UnallocatedPretaxAmount = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetUnallocatedPretaxAmountLocal(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.UnallocatedPretaxAmountLocal = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetUnallocatedDeductedByCashCoupons(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.UnallocatedDeductedByCashCoupons = &v
  return s
}

func (s *QueryInstanceGaapCostResponseDataModulesModule) SetUnallocatedDeductedByPrepaidCard(v string) *QueryInstanceGaapCostResponseDataModulesModule {
  s.UnallocatedDeductedByPrepaidCard = &v
  return s
}

type RenewInstanceRequest struct {
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  InstanceId *string `json:"InstanceId" xml:"InstanceId" require:"true"`
  RenewPeriod *int `json:"RenewPeriod" xml:"RenewPeriod" require:"true"`
  ClientToken *string `json:"ClientToken" xml:"ClientToken"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
}

func (s RenewInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
  return s.String()
}

func (s *RenewInstanceRequest) SetProductCode(v string) *RenewInstanceRequest {
  s.ProductCode = &v
  return s
}

func (s *RenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequest {
  s.InstanceId = &v
  return s
}

func (s *RenewInstanceRequest) SetRenewPeriod(v int) *RenewInstanceRequest {
  s.RenewPeriod = &v
  return s
}

func (s *RenewInstanceRequest) SetClientToken(v string) *RenewInstanceRequest {
  s.ClientToken = &v
  return s
}

func (s *RenewInstanceRequest) SetProductType(v string) *RenewInstanceRequest {
  s.ProductType = &v
  return s
}

func (s *RenewInstanceRequest) SetOwnerId(v int64) *RenewInstanceRequest {
  s.OwnerId = &v
  return s
}

type RenewInstanceResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *RenewInstanceResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s RenewInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s RenewInstanceResponse) GoString() string {
  return s.String()
}

func (s *RenewInstanceResponse) SetRequestId(v string) *RenewInstanceResponse {
  s.RequestId = &v
  return s
}

func (s *RenewInstanceResponse) SetSuccess(v bool) *RenewInstanceResponse {
  s.Success = &v
  return s
}

func (s *RenewInstanceResponse) SetCode(v string) *RenewInstanceResponse {
  s.Code = &v
  return s
}

func (s *RenewInstanceResponse) SetMessage(v string) *RenewInstanceResponse {
  s.Message = &v
  return s
}

func (s *RenewInstanceResponse) SetData(v *RenewInstanceResponseData) *RenewInstanceResponse {
  s.Data = v
  return s
}

type RenewInstanceResponseData struct {
  OrderId *string `json:"OrderId" xml:"OrderId" require:"true"`
}

func (s RenewInstanceResponseData) String() string {
  return tea.Prettify(s)
}

func (s RenewInstanceResponseData) GoString() string {
  return s.String()
}

func (s *RenewInstanceResponseData) SetOrderId(v string) *RenewInstanceResponseData {
  s.OrderId = &v
  return s
}

type GetOrderDetailRequest struct {
  OrderId *string `json:"OrderId" xml:"OrderId" require:"true"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
}

func (s GetOrderDetailRequest) String() string {
  return tea.Prettify(s)
}

func (s GetOrderDetailRequest) GoString() string {
  return s.String()
}

func (s *GetOrderDetailRequest) SetOrderId(v string) *GetOrderDetailRequest {
  s.OrderId = &v
  return s
}

func (s *GetOrderDetailRequest) SetOwnerId(v int64) *GetOrderDetailRequest {
  s.OwnerId = &v
  return s
}

type GetOrderDetailResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *GetOrderDetailResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s GetOrderDetailResponse) String() string {
  return tea.Prettify(s)
}

func (s GetOrderDetailResponse) GoString() string {
  return s.String()
}

func (s *GetOrderDetailResponse) SetRequestId(v string) *GetOrderDetailResponse {
  s.RequestId = &v
  return s
}

func (s *GetOrderDetailResponse) SetSuccess(v bool) *GetOrderDetailResponse {
  s.Success = &v
  return s
}

func (s *GetOrderDetailResponse) SetCode(v string) *GetOrderDetailResponse {
  s.Code = &v
  return s
}

func (s *GetOrderDetailResponse) SetMessage(v string) *GetOrderDetailResponse {
  s.Message = &v
  return s
}

func (s *GetOrderDetailResponse) SetData(v *GetOrderDetailResponseData) *GetOrderDetailResponse {
  s.Data = v
  return s
}

type GetOrderDetailResponseData struct {
  HostName *string `json:"HostName" xml:"HostName" require:"true"`
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  TotalCount *int `json:"TotalCount" xml:"TotalCount" require:"true"`
  OrderList *GetOrderDetailResponseDataOrderList `json:"OrderList" xml:"OrderList" require:"true" type:"Struct"`
}

func (s GetOrderDetailResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetOrderDetailResponseData) GoString() string {
  return s.String()
}

func (s *GetOrderDetailResponseData) SetHostName(v string) *GetOrderDetailResponseData {
  s.HostName = &v
  return s
}

func (s *GetOrderDetailResponseData) SetPageNum(v int) *GetOrderDetailResponseData {
  s.PageNum = &v
  return s
}

func (s *GetOrderDetailResponseData) SetPageSize(v int) *GetOrderDetailResponseData {
  s.PageSize = &v
  return s
}

func (s *GetOrderDetailResponseData) SetTotalCount(v int) *GetOrderDetailResponseData {
  s.TotalCount = &v
  return s
}

func (s *GetOrderDetailResponseData) SetOrderList(v *GetOrderDetailResponseDataOrderList) *GetOrderDetailResponseData {
  s.OrderList = v
  return s
}

type GetOrderDetailResponseDataOrderList struct {
  Order []*GetOrderDetailResponseDataOrderListOrder `json:"Order" xml:"Order" require:"true" type:"Repeated"`
}

func (s GetOrderDetailResponseDataOrderList) String() string {
  return tea.Prettify(s)
}

func (s GetOrderDetailResponseDataOrderList) GoString() string {
  return s.String()
}

func (s *GetOrderDetailResponseDataOrderList) SetOrder(v []*GetOrderDetailResponseDataOrderListOrder) *GetOrderDetailResponseDataOrderList {
  s.Order = v
  return s
}

type GetOrderDetailResponseDataOrderListOrder struct     {
  OrderId *string `json:"OrderId" xml:"OrderId" require:"true"`
  SubOrderId *string `json:"SubOrderId" xml:"SubOrderId" require:"true"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType" require:"true"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  OrderType *string `json:"OrderType" xml:"OrderType" require:"true"`
  CreateTime *string `json:"CreateTime" xml:"CreateTime" require:"true"`
  PaymentTime *string `json:"PaymentTime" xml:"PaymentTime" require:"true"`
  PaymentStatus *string `json:"PaymentStatus" xml:"PaymentStatus" require:"true"`
  Region *string `json:"Region" xml:"Region" require:"true"`
  Config *string `json:"Config" xml:"Config" require:"true"`
  Quantity *string `json:"Quantity" xml:"Quantity" require:"true"`
  UsageStartTime *string `json:"UsageStartTime" xml:"UsageStartTime" require:"true"`
  UsageEndTime *string `json:"UsageEndTime" xml:"UsageEndTime" require:"true"`
  InstanceIDs *string `json:"InstanceIDs" xml:"InstanceIDs" require:"true"`
  PretaxGrossAmount *string `json:"PretaxGrossAmount" xml:"PretaxGrossAmount" require:"true"`
  PretaxAmount *string `json:"PretaxAmount" xml:"PretaxAmount" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  PretaxAmountLocal *string `json:"PretaxAmountLocal" xml:"PretaxAmountLocal" require:"true"`
  Tax *string `json:"Tax" xml:"Tax" require:"true"`
  AfterTaxAmount *string `json:"AfterTaxAmount" xml:"AfterTaxAmount" require:"true"`
  PaymentCurrency *string `json:"PaymentCurrency" xml:"PaymentCurrency" require:"true"`
  Operator *string `json:"Operator" xml:"Operator" require:"true"`
  RelatedOrderId *string `json:"RelatedOrderId" xml:"RelatedOrderId" require:"true"`
  OrderSubType *string `json:"OrderSubType" xml:"OrderSubType" require:"true"`
  OriginalConfig *string `json:"OriginalConfig" xml:"OriginalConfig" require:"true"`
}

func (s GetOrderDetailResponseDataOrderListOrder) String() string {
  return tea.Prettify(s)
}

func (s GetOrderDetailResponseDataOrderListOrder) GoString() string {
  return s.String()
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetOrderId(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.OrderId = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetSubOrderId(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.SubOrderId = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetProductCode(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.ProductCode = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetProductType(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.ProductType = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetSubscriptionType(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.SubscriptionType = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetOrderType(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.OrderType = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetCreateTime(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.CreateTime = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetPaymentTime(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.PaymentTime = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetPaymentStatus(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.PaymentStatus = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetRegion(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.Region = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetConfig(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.Config = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetQuantity(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.Quantity = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetUsageStartTime(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.UsageStartTime = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetUsageEndTime(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.UsageEndTime = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetInstanceIDs(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.InstanceIDs = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetPretaxGrossAmount(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.PretaxGrossAmount = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetPretaxAmount(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.PretaxAmount = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetCurrency(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.Currency = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetPretaxAmountLocal(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.PretaxAmountLocal = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetTax(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.Tax = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetAfterTaxAmount(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.AfterTaxAmount = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetPaymentCurrency(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.PaymentCurrency = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetOperator(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.Operator = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetRelatedOrderId(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.RelatedOrderId = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetOrderSubType(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.OrderSubType = &v
  return s
}

func (s *GetOrderDetailResponseDataOrderListOrder) SetOriginalConfig(v string) *GetOrderDetailResponseDataOrderListOrder {
  s.OriginalConfig = &v
  return s
}

type QueryOrdersRequest struct {
  CreateTimeEnd *string `json:"CreateTimeEnd" xml:"CreateTimeEnd"`
  PageNum *int `json:"PageNum" xml:"PageNum"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType"`
  OrderType *string `json:"OrderType" xml:"OrderType"`
  PaymentStatus *string `json:"PaymentStatus" xml:"PaymentStatus"`
  CreateTimeStart *string `json:"CreateTimeStart" xml:"CreateTimeStart"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
}

func (s QueryOrdersRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryOrdersRequest) GoString() string {
  return s.String()
}

func (s *QueryOrdersRequest) SetCreateTimeEnd(v string) *QueryOrdersRequest {
  s.CreateTimeEnd = &v
  return s
}

func (s *QueryOrdersRequest) SetPageNum(v int) *QueryOrdersRequest {
  s.PageNum = &v
  return s
}

func (s *QueryOrdersRequest) SetPageSize(v int) *QueryOrdersRequest {
  s.PageSize = &v
  return s
}

func (s *QueryOrdersRequest) SetProductCode(v string) *QueryOrdersRequest {
  s.ProductCode = &v
  return s
}

func (s *QueryOrdersRequest) SetProductType(v string) *QueryOrdersRequest {
  s.ProductType = &v
  return s
}

func (s *QueryOrdersRequest) SetSubscriptionType(v string) *QueryOrdersRequest {
  s.SubscriptionType = &v
  return s
}

func (s *QueryOrdersRequest) SetOrderType(v string) *QueryOrdersRequest {
  s.OrderType = &v
  return s
}

func (s *QueryOrdersRequest) SetPaymentStatus(v string) *QueryOrdersRequest {
  s.PaymentStatus = &v
  return s
}

func (s *QueryOrdersRequest) SetCreateTimeStart(v string) *QueryOrdersRequest {
  s.CreateTimeStart = &v
  return s
}

func (s *QueryOrdersRequest) SetOwnerId(v int64) *QueryOrdersRequest {
  s.OwnerId = &v
  return s
}

type QueryOrdersResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryOrdersResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryOrdersResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOrdersResponse) GoString() string {
  return s.String()
}

func (s *QueryOrdersResponse) SetRequestId(v string) *QueryOrdersResponse {
  s.RequestId = &v
  return s
}

func (s *QueryOrdersResponse) SetSuccess(v bool) *QueryOrdersResponse {
  s.Success = &v
  return s
}

func (s *QueryOrdersResponse) SetCode(v string) *QueryOrdersResponse {
  s.Code = &v
  return s
}

func (s *QueryOrdersResponse) SetMessage(v string) *QueryOrdersResponse {
  s.Message = &v
  return s
}

func (s *QueryOrdersResponse) SetData(v *QueryOrdersResponseData) *QueryOrdersResponse {
  s.Data = v
  return s
}

type QueryOrdersResponseData struct {
  HostName *string `json:"HostName" xml:"HostName" require:"true"`
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  TotalCount *int `json:"TotalCount" xml:"TotalCount" require:"true"`
  OrderList *QueryOrdersResponseDataOrderList `json:"OrderList" xml:"OrderList" require:"true" type:"Struct"`
}

func (s QueryOrdersResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryOrdersResponseData) GoString() string {
  return s.String()
}

func (s *QueryOrdersResponseData) SetHostName(v string) *QueryOrdersResponseData {
  s.HostName = &v
  return s
}

func (s *QueryOrdersResponseData) SetPageNum(v int) *QueryOrdersResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryOrdersResponseData) SetPageSize(v int) *QueryOrdersResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryOrdersResponseData) SetTotalCount(v int) *QueryOrdersResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryOrdersResponseData) SetOrderList(v *QueryOrdersResponseDataOrderList) *QueryOrdersResponseData {
  s.OrderList = v
  return s
}

type QueryOrdersResponseDataOrderList struct {
  Order []*QueryOrdersResponseDataOrderListOrder `json:"Order" xml:"Order" require:"true" type:"Repeated"`
}

func (s QueryOrdersResponseDataOrderList) String() string {
  return tea.Prettify(s)
}

func (s QueryOrdersResponseDataOrderList) GoString() string {
  return s.String()
}

func (s *QueryOrdersResponseDataOrderList) SetOrder(v []*QueryOrdersResponseDataOrderListOrder) *QueryOrdersResponseDataOrderList {
  s.Order = v
  return s
}

type QueryOrdersResponseDataOrderListOrder struct     {
  OrderId *string `json:"OrderId" xml:"OrderId" require:"true"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType" require:"true"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  OrderType *string `json:"OrderType" xml:"OrderType" require:"true"`
  CreateTime *string `json:"CreateTime" xml:"CreateTime" require:"true"`
  PaymentTime *string `json:"PaymentTime" xml:"PaymentTime" require:"true"`
  PaymentStatus *string `json:"PaymentStatus" xml:"PaymentStatus" require:"true"`
  PretaxGrossAmount *string `json:"PretaxGrossAmount" xml:"PretaxGrossAmount" require:"true"`
  PretaxAmount *string `json:"PretaxAmount" xml:"PretaxAmount" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  PretaxAmountLocal *string `json:"PretaxAmountLocal" xml:"PretaxAmountLocal" require:"true"`
  Tax *string `json:"Tax" xml:"Tax" require:"true"`
  AfterTaxAmount *string `json:"AfterTaxAmount" xml:"AfterTaxAmount" require:"true"`
  PaymentCurrency *string `json:"PaymentCurrency" xml:"PaymentCurrency" require:"true"`
  RelatedOrderId *string `json:"RelatedOrderId" xml:"RelatedOrderId" require:"true"`
}

func (s QueryOrdersResponseDataOrderListOrder) String() string {
  return tea.Prettify(s)
}

func (s QueryOrdersResponseDataOrderListOrder) GoString() string {
  return s.String()
}

func (s *QueryOrdersResponseDataOrderListOrder) SetOrderId(v string) *QueryOrdersResponseDataOrderListOrder {
  s.OrderId = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetProductCode(v string) *QueryOrdersResponseDataOrderListOrder {
  s.ProductCode = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetProductType(v string) *QueryOrdersResponseDataOrderListOrder {
  s.ProductType = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetSubscriptionType(v string) *QueryOrdersResponseDataOrderListOrder {
  s.SubscriptionType = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetOrderType(v string) *QueryOrdersResponseDataOrderListOrder {
  s.OrderType = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetCreateTime(v string) *QueryOrdersResponseDataOrderListOrder {
  s.CreateTime = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetPaymentTime(v string) *QueryOrdersResponseDataOrderListOrder {
  s.PaymentTime = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetPaymentStatus(v string) *QueryOrdersResponseDataOrderListOrder {
  s.PaymentStatus = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetPretaxGrossAmount(v string) *QueryOrdersResponseDataOrderListOrder {
  s.PretaxGrossAmount = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetPretaxAmount(v string) *QueryOrdersResponseDataOrderListOrder {
  s.PretaxAmount = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetCurrency(v string) *QueryOrdersResponseDataOrderListOrder {
  s.Currency = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetPretaxAmountLocal(v string) *QueryOrdersResponseDataOrderListOrder {
  s.PretaxAmountLocal = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetTax(v string) *QueryOrdersResponseDataOrderListOrder {
  s.Tax = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetAfterTaxAmount(v string) *QueryOrdersResponseDataOrderListOrder {
  s.AfterTaxAmount = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetPaymentCurrency(v string) *QueryOrdersResponseDataOrderListOrder {
  s.PaymentCurrency = &v
  return s
}

func (s *QueryOrdersResponseDataOrderListOrder) SetRelatedOrderId(v string) *QueryOrdersResponseDataOrderListOrder {
  s.RelatedOrderId = &v
  return s
}

type QueryMonthlyInstanceConsumptionRequest struct {
  ProductCode *string `json:"ProductCode" xml:"ProductCode"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  PageNum *int `json:"PageNum" xml:"PageNum"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType"`
}

func (s QueryMonthlyInstanceConsumptionRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryMonthlyInstanceConsumptionRequest) GoString() string {
  return s.String()
}

func (s *QueryMonthlyInstanceConsumptionRequest) SetProductCode(v string) *QueryMonthlyInstanceConsumptionRequest {
  s.ProductCode = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionRequest) SetOwnerId(v int64) *QueryMonthlyInstanceConsumptionRequest {
  s.OwnerId = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionRequest) SetPageNum(v int) *QueryMonthlyInstanceConsumptionRequest {
  s.PageNum = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionRequest) SetPageSize(v int) *QueryMonthlyInstanceConsumptionRequest {
  s.PageSize = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionRequest) SetBillingCycle(v string) *QueryMonthlyInstanceConsumptionRequest {
  s.BillingCycle = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionRequest) SetProductType(v string) *QueryMonthlyInstanceConsumptionRequest {
  s.ProductType = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionRequest) SetSubscriptionType(v string) *QueryMonthlyInstanceConsumptionRequest {
  s.SubscriptionType = &v
  return s
}

type QueryMonthlyInstanceConsumptionResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryMonthlyInstanceConsumptionResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryMonthlyInstanceConsumptionResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryMonthlyInstanceConsumptionResponse) GoString() string {
  return s.String()
}

func (s *QueryMonthlyInstanceConsumptionResponse) SetRequestId(v string) *QueryMonthlyInstanceConsumptionResponse {
  s.RequestId = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponse) SetSuccess(v bool) *QueryMonthlyInstanceConsumptionResponse {
  s.Success = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponse) SetCode(v string) *QueryMonthlyInstanceConsumptionResponse {
  s.Code = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponse) SetMessage(v string) *QueryMonthlyInstanceConsumptionResponse {
  s.Message = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponse) SetData(v *QueryMonthlyInstanceConsumptionResponseData) *QueryMonthlyInstanceConsumptionResponse {
  s.Data = v
  return s
}

type QueryMonthlyInstanceConsumptionResponseData struct {
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  TotalCount *int `json:"TotalCount" xml:"TotalCount" require:"true"`
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  Items *QueryMonthlyInstanceConsumptionResponseDataItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s QueryMonthlyInstanceConsumptionResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryMonthlyInstanceConsumptionResponseData) GoString() string {
  return s.String()
}

func (s *QueryMonthlyInstanceConsumptionResponseData) SetPageNum(v int) *QueryMonthlyInstanceConsumptionResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseData) SetPageSize(v int) *QueryMonthlyInstanceConsumptionResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseData) SetTotalCount(v int) *QueryMonthlyInstanceConsumptionResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseData) SetBillingCycle(v string) *QueryMonthlyInstanceConsumptionResponseData {
  s.BillingCycle = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseData) SetItems(v *QueryMonthlyInstanceConsumptionResponseDataItems) *QueryMonthlyInstanceConsumptionResponseData {
  s.Items = v
  return s
}

type QueryMonthlyInstanceConsumptionResponseDataItems struct {
  Item []*QueryMonthlyInstanceConsumptionResponseDataItemsItem `json:"Item" xml:"Item" require:"true" type:"Repeated"`
}

func (s QueryMonthlyInstanceConsumptionResponseDataItems) String() string {
  return tea.Prettify(s)
}

func (s QueryMonthlyInstanceConsumptionResponseDataItems) GoString() string {
  return s.String()
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItems) SetItem(v []*QueryMonthlyInstanceConsumptionResponseDataItemsItem) *QueryMonthlyInstanceConsumptionResponseDataItems {
  s.Item = v
  return s
}

type QueryMonthlyInstanceConsumptionResponseDataItemsItem struct     {
  InstanceID *string `json:"InstanceID" xml:"InstanceID" require:"true"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType" require:"true"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  Tag *string `json:"Tag" xml:"Tag" require:"true"`
  ResourceGroup *string `json:"ResourceGroup" xml:"ResourceGroup" require:"true"`
  PayerAccount *string `json:"PayerAccount" xml:"PayerAccount" require:"true"`
  OwnerID *string `json:"OwnerID" xml:"OwnerID" require:"true"`
  Region *string `json:"Region" xml:"Region" require:"true"`
  PretaxGrossAmount *float32 `json:"PretaxGrossAmount" xml:"PretaxGrossAmount" require:"true"`
  DiscountAmount *float32 `json:"DiscountAmount" xml:"DiscountAmount" require:"true"`
  PretaxAmount *float32 `json:"PretaxAmount" xml:"PretaxAmount" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  PretaxAmountLocal *float32 `json:"PretaxAmountLocal" xml:"PretaxAmountLocal" require:"true"`
  Tax *float32 `json:"Tax" xml:"Tax" require:"true"`
  AfterTaxAmount *float32 `json:"AfterTaxAmount" xml:"AfterTaxAmount" require:"true"`
  PaymentCurrency *string `json:"PaymentCurrency" xml:"PaymentCurrency" require:"true"`
}

func (s QueryMonthlyInstanceConsumptionResponseDataItemsItem) String() string {
  return tea.Prettify(s)
}

func (s QueryMonthlyInstanceConsumptionResponseDataItemsItem) GoString() string {
  return s.String()
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetInstanceID(v string) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.InstanceID = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetProductCode(v string) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.ProductCode = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetProductType(v string) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.ProductType = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetSubscriptionType(v string) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.SubscriptionType = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetTag(v string) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.Tag = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetResourceGroup(v string) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.ResourceGroup = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetPayerAccount(v string) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.PayerAccount = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetOwnerID(v string) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.OwnerID = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetRegion(v string) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.Region = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetPretaxGrossAmount(v float32) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.PretaxGrossAmount = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetDiscountAmount(v float32) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.DiscountAmount = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetPretaxAmount(v float32) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.PretaxAmount = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetCurrency(v string) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.Currency = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetPretaxAmountLocal(v float32) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.PretaxAmountLocal = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetTax(v float32) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.Tax = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetAfterTaxAmount(v float32) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.AfterTaxAmount = &v
  return s
}

func (s *QueryMonthlyInstanceConsumptionResponseDataItemsItem) SetPaymentCurrency(v string) *QueryMonthlyInstanceConsumptionResponseDataItemsItem {
  s.PaymentCurrency = &v
  return s
}

type QuerySettlementBillRequest struct {
  PageSize *int `json:"PageSize" xml:"PageSize"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  PageNum *int `json:"PageNum" xml:"PageNum"`
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  StartTime *string `json:"StartTime" xml:"StartTime"`
  EndTime *string `json:"EndTime" xml:"EndTime"`
  Type *string `json:"Type" xml:"Type"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType"`
  IsHideZeroCharge *bool `json:"IsHideZeroCharge" xml:"IsHideZeroCharge"`
}

func (s QuerySettlementBillRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerySettlementBillRequest) GoString() string {
  return s.String()
}

func (s *QuerySettlementBillRequest) SetPageSize(v int) *QuerySettlementBillRequest {
  s.PageSize = &v
  return s
}

func (s *QuerySettlementBillRequest) SetOwnerId(v int64) *QuerySettlementBillRequest {
  s.OwnerId = &v
  return s
}

func (s *QuerySettlementBillRequest) SetPageNum(v int) *QuerySettlementBillRequest {
  s.PageNum = &v
  return s
}

func (s *QuerySettlementBillRequest) SetBillingCycle(v string) *QuerySettlementBillRequest {
  s.BillingCycle = &v
  return s
}

func (s *QuerySettlementBillRequest) SetStartTime(v string) *QuerySettlementBillRequest {
  s.StartTime = &v
  return s
}

func (s *QuerySettlementBillRequest) SetEndTime(v string) *QuerySettlementBillRequest {
  s.EndTime = &v
  return s
}

func (s *QuerySettlementBillRequest) SetType(v string) *QuerySettlementBillRequest {
  s.Type = &v
  return s
}

func (s *QuerySettlementBillRequest) SetProductCode(v string) *QuerySettlementBillRequest {
  s.ProductCode = &v
  return s
}

func (s *QuerySettlementBillRequest) SetProductType(v string) *QuerySettlementBillRequest {
  s.ProductType = &v
  return s
}

func (s *QuerySettlementBillRequest) SetSubscriptionType(v string) *QuerySettlementBillRequest {
  s.SubscriptionType = &v
  return s
}

func (s *QuerySettlementBillRequest) SetIsHideZeroCharge(v bool) *QuerySettlementBillRequest {
  s.IsHideZeroCharge = &v
  return s
}

type QuerySettlementBillResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QuerySettlementBillResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QuerySettlementBillResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerySettlementBillResponse) GoString() string {
  return s.String()
}

func (s *QuerySettlementBillResponse) SetRequestId(v string) *QuerySettlementBillResponse {
  s.RequestId = &v
  return s
}

func (s *QuerySettlementBillResponse) SetSuccess(v bool) *QuerySettlementBillResponse {
  s.Success = &v
  return s
}

func (s *QuerySettlementBillResponse) SetCode(v string) *QuerySettlementBillResponse {
  s.Code = &v
  return s
}

func (s *QuerySettlementBillResponse) SetMessage(v string) *QuerySettlementBillResponse {
  s.Message = &v
  return s
}

func (s *QuerySettlementBillResponse) SetData(v *QuerySettlementBillResponseData) *QuerySettlementBillResponse {
  s.Data = v
  return s
}

type QuerySettlementBillResponseData struct {
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  TotalCount *int `json:"TotalCount" xml:"TotalCount" require:"true"`
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  Items *QuerySettlementBillResponseDataItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s QuerySettlementBillResponseData) String() string {
  return tea.Prettify(s)
}

func (s QuerySettlementBillResponseData) GoString() string {
  return s.String()
}

func (s *QuerySettlementBillResponseData) SetPageNum(v int) *QuerySettlementBillResponseData {
  s.PageNum = &v
  return s
}

func (s *QuerySettlementBillResponseData) SetPageSize(v int) *QuerySettlementBillResponseData {
  s.PageSize = &v
  return s
}

func (s *QuerySettlementBillResponseData) SetTotalCount(v int) *QuerySettlementBillResponseData {
  s.TotalCount = &v
  return s
}

func (s *QuerySettlementBillResponseData) SetBillingCycle(v string) *QuerySettlementBillResponseData {
  s.BillingCycle = &v
  return s
}

func (s *QuerySettlementBillResponseData) SetItems(v *QuerySettlementBillResponseDataItems) *QuerySettlementBillResponseData {
  s.Items = v
  return s
}

type QuerySettlementBillResponseDataItems struct {
  Item []*QuerySettlementBillResponseDataItemsItem `json:"Item" xml:"Item" require:"true" type:"Repeated"`
}

func (s QuerySettlementBillResponseDataItems) String() string {
  return tea.Prettify(s)
}

func (s QuerySettlementBillResponseDataItems) GoString() string {
  return s.String()
}

func (s *QuerySettlementBillResponseDataItems) SetItem(v []*QuerySettlementBillResponseDataItemsItem) *QuerySettlementBillResponseDataItems {
  s.Item = v
  return s
}

type QuerySettlementBillResponseDataItemsItem struct     {
  RecordID *string `json:"RecordID" xml:"RecordID" require:"true"`
  Item *string `json:"Item" xml:"Item" require:"true"`
  PayerAccount *string `json:"PayerAccount" xml:"PayerAccount" require:"true"`
  OwnerID *string `json:"OwnerID" xml:"OwnerID" require:"true"`
  CreateTime *string `json:"CreateTime" xml:"CreateTime" require:"true"`
  UsageStartTime *string `json:"UsageStartTime" xml:"UsageStartTime" require:"true"`
  UsageEndTime *string `json:"UsageEndTime" xml:"UsageEndTime" require:"true"`
  SuborderID *string `json:"SuborderID" xml:"SuborderID" require:"true"`
  OrderID *string `json:"OrderID" xml:"OrderID" require:"true"`
  OrderType *string `json:"OrderType" xml:"OrderType" require:"true"`
  LinkedCustomerOrderID *string `json:"LinkedCustomerOrderID" xml:"LinkedCustomerOrderID" require:"true"`
  OriginalOrderID *string `json:"OriginalOrderID" xml:"OriginalOrderID" require:"true"`
  PaymentTime *string `json:"PaymentTime" xml:"PaymentTime" require:"true"`
  SolutionID *string `json:"SolutionID" xml:"SolutionID" require:"true"`
  SolutionName *string `json:"SolutionName" xml:"SolutionName" require:"true"`
  BillID *string `json:"BillID" xml:"BillID" require:"true"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType" require:"true"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  Region *string `json:"Region" xml:"Region" require:"true"`
  Config *string `json:"Config" xml:"Config" require:"true"`
  Quantity *string `json:"Quantity" xml:"Quantity" require:"true"`
  PretaxGrossAmount *float32 `json:"PretaxGrossAmount" xml:"PretaxGrossAmount" require:"true"`
  ChargeDiscount *float32 `json:"ChargeDiscount" xml:"ChargeDiscount" require:"true"`
  DeductedByCoupons *float32 `json:"DeductedByCoupons" xml:"DeductedByCoupons" require:"true"`
  AccountDiscount *float32 `json:"AccountDiscount" xml:"AccountDiscount" require:"true"`
  Promotion *string `json:"Promotion" xml:"Promotion" require:"true"`
  PretaxAmount *float32 `json:"PretaxAmount" xml:"PretaxAmount" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  PretaxAmountLocal *float32 `json:"PretaxAmountLocal" xml:"PretaxAmountLocal" require:"true"`
  PreviousBillingCycleBalance *float32 `json:"PreviousBillingCycleBalance" xml:"PreviousBillingCycleBalance" require:"true"`
  Tax *float32 `json:"Tax" xml:"Tax" require:"true"`
  AfterTaxAmount *float32 `json:"AfterTaxAmount" xml:"AfterTaxAmount" require:"true"`
  Status *string `json:"Status" xml:"Status" require:"true"`
  ClearedTime *string `json:"ClearedTime" xml:"ClearedTime" require:"true"`
  OutstandingAmount *float32 `json:"OutstandingAmount" xml:"OutstandingAmount" require:"true"`
  DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons" xml:"DeductedByCashCoupons" require:"true"`
  DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard" xml:"DeductedByPrepaidCard" require:"true"`
  MybankPaymentAmount *float32 `json:"MybankPaymentAmount" xml:"MybankPaymentAmount" require:"true"`
  PaymentAmount *float32 `json:"PaymentAmount" xml:"PaymentAmount" require:"true"`
  PaymentCurrency *string `json:"PaymentCurrency" xml:"PaymentCurrency" require:"true"`
  Seller *string `json:"Seller" xml:"Seller" require:"true"`
  InvoiceNo *string `json:"InvoiceNo" xml:"InvoiceNo" require:"true"`
}

func (s QuerySettlementBillResponseDataItemsItem) String() string {
  return tea.Prettify(s)
}

func (s QuerySettlementBillResponseDataItemsItem) GoString() string {
  return s.String()
}

func (s *QuerySettlementBillResponseDataItemsItem) SetRecordID(v string) *QuerySettlementBillResponseDataItemsItem {
  s.RecordID = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetItem(v string) *QuerySettlementBillResponseDataItemsItem {
  s.Item = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetPayerAccount(v string) *QuerySettlementBillResponseDataItemsItem {
  s.PayerAccount = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetOwnerID(v string) *QuerySettlementBillResponseDataItemsItem {
  s.OwnerID = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetCreateTime(v string) *QuerySettlementBillResponseDataItemsItem {
  s.CreateTime = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetUsageStartTime(v string) *QuerySettlementBillResponseDataItemsItem {
  s.UsageStartTime = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetUsageEndTime(v string) *QuerySettlementBillResponseDataItemsItem {
  s.UsageEndTime = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetSuborderID(v string) *QuerySettlementBillResponseDataItemsItem {
  s.SuborderID = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetOrderID(v string) *QuerySettlementBillResponseDataItemsItem {
  s.OrderID = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetOrderType(v string) *QuerySettlementBillResponseDataItemsItem {
  s.OrderType = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetLinkedCustomerOrderID(v string) *QuerySettlementBillResponseDataItemsItem {
  s.LinkedCustomerOrderID = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetOriginalOrderID(v string) *QuerySettlementBillResponseDataItemsItem {
  s.OriginalOrderID = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetPaymentTime(v string) *QuerySettlementBillResponseDataItemsItem {
  s.PaymentTime = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetSolutionID(v string) *QuerySettlementBillResponseDataItemsItem {
  s.SolutionID = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetSolutionName(v string) *QuerySettlementBillResponseDataItemsItem {
  s.SolutionName = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetBillID(v string) *QuerySettlementBillResponseDataItemsItem {
  s.BillID = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetProductCode(v string) *QuerySettlementBillResponseDataItemsItem {
  s.ProductCode = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetProductType(v string) *QuerySettlementBillResponseDataItemsItem {
  s.ProductType = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetSubscriptionType(v string) *QuerySettlementBillResponseDataItemsItem {
  s.SubscriptionType = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetRegion(v string) *QuerySettlementBillResponseDataItemsItem {
  s.Region = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetConfig(v string) *QuerySettlementBillResponseDataItemsItem {
  s.Config = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetQuantity(v string) *QuerySettlementBillResponseDataItemsItem {
  s.Quantity = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetPretaxGrossAmount(v float32) *QuerySettlementBillResponseDataItemsItem {
  s.PretaxGrossAmount = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetChargeDiscount(v float32) *QuerySettlementBillResponseDataItemsItem {
  s.ChargeDiscount = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetDeductedByCoupons(v float32) *QuerySettlementBillResponseDataItemsItem {
  s.DeductedByCoupons = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetAccountDiscount(v float32) *QuerySettlementBillResponseDataItemsItem {
  s.AccountDiscount = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetPromotion(v string) *QuerySettlementBillResponseDataItemsItem {
  s.Promotion = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetPretaxAmount(v float32) *QuerySettlementBillResponseDataItemsItem {
  s.PretaxAmount = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetCurrency(v string) *QuerySettlementBillResponseDataItemsItem {
  s.Currency = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetPretaxAmountLocal(v float32) *QuerySettlementBillResponseDataItemsItem {
  s.PretaxAmountLocal = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetPreviousBillingCycleBalance(v float32) *QuerySettlementBillResponseDataItemsItem {
  s.PreviousBillingCycleBalance = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetTax(v float32) *QuerySettlementBillResponseDataItemsItem {
  s.Tax = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetAfterTaxAmount(v float32) *QuerySettlementBillResponseDataItemsItem {
  s.AfterTaxAmount = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetStatus(v string) *QuerySettlementBillResponseDataItemsItem {
  s.Status = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetClearedTime(v string) *QuerySettlementBillResponseDataItemsItem {
  s.ClearedTime = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetOutstandingAmount(v float32) *QuerySettlementBillResponseDataItemsItem {
  s.OutstandingAmount = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetDeductedByCashCoupons(v float32) *QuerySettlementBillResponseDataItemsItem {
  s.DeductedByCashCoupons = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetDeductedByPrepaidCard(v float32) *QuerySettlementBillResponseDataItemsItem {
  s.DeductedByPrepaidCard = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetMybankPaymentAmount(v float32) *QuerySettlementBillResponseDataItemsItem {
  s.MybankPaymentAmount = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetPaymentAmount(v float32) *QuerySettlementBillResponseDataItemsItem {
  s.PaymentAmount = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetPaymentCurrency(v string) *QuerySettlementBillResponseDataItemsItem {
  s.PaymentCurrency = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetSeller(v string) *QuerySettlementBillResponseDataItemsItem {
  s.Seller = &v
  return s
}

func (s *QuerySettlementBillResponseDataItemsItem) SetInvoiceNo(v string) *QuerySettlementBillResponseDataItemsItem {
  s.InvoiceNo = &v
  return s
}

type QueryMonthlyBillRequest struct {
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
}

func (s QueryMonthlyBillRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryMonthlyBillRequest) GoString() string {
  return s.String()
}

func (s *QueryMonthlyBillRequest) SetBillingCycle(v string) *QueryMonthlyBillRequest {
  s.BillingCycle = &v
  return s
}

type QueryMonthlyBillResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryMonthlyBillResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryMonthlyBillResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryMonthlyBillResponse) GoString() string {
  return s.String()
}

func (s *QueryMonthlyBillResponse) SetRequestId(v string) *QueryMonthlyBillResponse {
  s.RequestId = &v
  return s
}

func (s *QueryMonthlyBillResponse) SetSuccess(v bool) *QueryMonthlyBillResponse {
  s.Success = &v
  return s
}

func (s *QueryMonthlyBillResponse) SetCode(v string) *QueryMonthlyBillResponse {
  s.Code = &v
  return s
}

func (s *QueryMonthlyBillResponse) SetMessage(v string) *QueryMonthlyBillResponse {
  s.Message = &v
  return s
}

func (s *QueryMonthlyBillResponse) SetData(v *QueryMonthlyBillResponseData) *QueryMonthlyBillResponse {
  s.Data = v
  return s
}

type QueryMonthlyBillResponseData struct {
  OutstandingAmount *float32 `json:"OutstandingAmount" xml:"OutstandingAmount" require:"true"`
  TotalOutstandingAmount *float32 `json:"TotalOutstandingAmount" xml:"TotalOutstandingAmount" require:"true"`
  NewInvoiceAmount *float32 `json:"NewInvoiceAmount" xml:"NewInvoiceAmount" require:"true"`
  BillingCycle *string `json:"BillingCycle" xml:"BillingCycle" require:"true"`
  Items *QueryMonthlyBillResponseDataItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s QueryMonthlyBillResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryMonthlyBillResponseData) GoString() string {
  return s.String()
}

func (s *QueryMonthlyBillResponseData) SetOutstandingAmount(v float32) *QueryMonthlyBillResponseData {
  s.OutstandingAmount = &v
  return s
}

func (s *QueryMonthlyBillResponseData) SetTotalOutstandingAmount(v float32) *QueryMonthlyBillResponseData {
  s.TotalOutstandingAmount = &v
  return s
}

func (s *QueryMonthlyBillResponseData) SetNewInvoiceAmount(v float32) *QueryMonthlyBillResponseData {
  s.NewInvoiceAmount = &v
  return s
}

func (s *QueryMonthlyBillResponseData) SetBillingCycle(v string) *QueryMonthlyBillResponseData {
  s.BillingCycle = &v
  return s
}

func (s *QueryMonthlyBillResponseData) SetItems(v *QueryMonthlyBillResponseDataItems) *QueryMonthlyBillResponseData {
  s.Items = v
  return s
}

type QueryMonthlyBillResponseDataItems struct {
  Item []*QueryMonthlyBillResponseDataItemsItem `json:"Item" xml:"Item" require:"true" type:"Repeated"`
}

func (s QueryMonthlyBillResponseDataItems) String() string {
  return tea.Prettify(s)
}

func (s QueryMonthlyBillResponseDataItems) GoString() string {
  return s.String()
}

func (s *QueryMonthlyBillResponseDataItems) SetItem(v []*QueryMonthlyBillResponseDataItemsItem) *QueryMonthlyBillResponseDataItems {
  s.Item = v
  return s
}

type QueryMonthlyBillResponseDataItemsItem struct     {
  Item *string `json:"Item" xml:"Item" require:"true"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType" require:"true"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  SolutionCode *string `json:"SolutionCode" xml:"SolutionCode" require:"true"`
  SolutionName *string `json:"SolutionName" xml:"SolutionName" require:"true"`
  PretaxGrossAmount *float32 `json:"PretaxGrossAmount" xml:"PretaxGrossAmount" require:"true"`
  InvoiceDiscount *float32 `json:"InvoiceDiscount" xml:"InvoiceDiscount" require:"true"`
  DeductedByCoupons *float32 `json:"DeductedByCoupons" xml:"DeductedByCoupons" require:"true"`
  PretaxAmount *float32 `json:"PretaxAmount" xml:"PretaxAmount" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  PretaxAmountLocal *float32 `json:"PretaxAmountLocal" xml:"PretaxAmountLocal" require:"true"`
  Tax *float32 `json:"Tax" xml:"Tax" require:"true"`
  AfterTaxAmount *float32 `json:"AfterTaxAmount" xml:"AfterTaxAmount" require:"true"`
  OutstandingAmount *float32 `json:"OutstandingAmount" xml:"OutstandingAmount" require:"true"`
  DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons" xml:"DeductedByCashCoupons" require:"true"`
  DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard" xml:"DeductedByPrepaidCard" require:"true"`
  PaymentAmount *float32 `json:"PaymentAmount" xml:"PaymentAmount" require:"true"`
  PaymentCurrency *string `json:"PaymentCurrency" xml:"PaymentCurrency" require:"true"`
}

func (s QueryMonthlyBillResponseDataItemsItem) String() string {
  return tea.Prettify(s)
}

func (s QueryMonthlyBillResponseDataItemsItem) GoString() string {
  return s.String()
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetItem(v string) *QueryMonthlyBillResponseDataItemsItem {
  s.Item = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetProductCode(v string) *QueryMonthlyBillResponseDataItemsItem {
  s.ProductCode = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetProductType(v string) *QueryMonthlyBillResponseDataItemsItem {
  s.ProductType = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetSubscriptionType(v string) *QueryMonthlyBillResponseDataItemsItem {
  s.SubscriptionType = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetSolutionCode(v string) *QueryMonthlyBillResponseDataItemsItem {
  s.SolutionCode = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetSolutionName(v string) *QueryMonthlyBillResponseDataItemsItem {
  s.SolutionName = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetPretaxGrossAmount(v float32) *QueryMonthlyBillResponseDataItemsItem {
  s.PretaxGrossAmount = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetInvoiceDiscount(v float32) *QueryMonthlyBillResponseDataItemsItem {
  s.InvoiceDiscount = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetDeductedByCoupons(v float32) *QueryMonthlyBillResponseDataItemsItem {
  s.DeductedByCoupons = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetPretaxAmount(v float32) *QueryMonthlyBillResponseDataItemsItem {
  s.PretaxAmount = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetCurrency(v string) *QueryMonthlyBillResponseDataItemsItem {
  s.Currency = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetPretaxAmountLocal(v float32) *QueryMonthlyBillResponseDataItemsItem {
  s.PretaxAmountLocal = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetTax(v float32) *QueryMonthlyBillResponseDataItemsItem {
  s.Tax = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetAfterTaxAmount(v float32) *QueryMonthlyBillResponseDataItemsItem {
  s.AfterTaxAmount = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetOutstandingAmount(v float32) *QueryMonthlyBillResponseDataItemsItem {
  s.OutstandingAmount = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetDeductedByCashCoupons(v float32) *QueryMonthlyBillResponseDataItemsItem {
  s.DeductedByCashCoupons = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetDeductedByPrepaidCard(v float32) *QueryMonthlyBillResponseDataItemsItem {
  s.DeductedByPrepaidCard = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetPaymentAmount(v float32) *QueryMonthlyBillResponseDataItemsItem {
  s.PaymentAmount = &v
  return s
}

func (s *QueryMonthlyBillResponseDataItemsItem) SetPaymentCurrency(v string) *QueryMonthlyBillResponseDataItemsItem {
  s.PaymentCurrency = &v
  return s
}

type SetRenewalRequest struct {
  RenewalPeriod *int `json:"RenewalPeriod" xml:"RenewalPeriod"`
  InstanceIDs *string `json:"InstanceIDs" xml:"InstanceIDs" require:"true"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType"`
  RenewalPeriodUnit *string `json:"RenewalPeriodUnit" xml:"RenewalPeriodUnit"`
  RenewalStatus *string `json:"RenewalStatus" xml:"RenewalStatus" require:"true"`
}

func (s SetRenewalRequest) String() string {
  return tea.Prettify(s)
}

func (s SetRenewalRequest) GoString() string {
  return s.String()
}

func (s *SetRenewalRequest) SetRenewalPeriod(v int) *SetRenewalRequest {
  s.RenewalPeriod = &v
  return s
}

func (s *SetRenewalRequest) SetInstanceIDs(v string) *SetRenewalRequest {
  s.InstanceIDs = &v
  return s
}

func (s *SetRenewalRequest) SetOwnerId(v int64) *SetRenewalRequest {
  s.OwnerId = &v
  return s
}

func (s *SetRenewalRequest) SetProductCode(v string) *SetRenewalRequest {
  s.ProductCode = &v
  return s
}

func (s *SetRenewalRequest) SetProductType(v string) *SetRenewalRequest {
  s.ProductType = &v
  return s
}

func (s *SetRenewalRequest) SetSubscriptionType(v string) *SetRenewalRequest {
  s.SubscriptionType = &v
  return s
}

func (s *SetRenewalRequest) SetRenewalPeriodUnit(v string) *SetRenewalRequest {
  s.RenewalPeriodUnit = &v
  return s
}

func (s *SetRenewalRequest) SetRenewalStatus(v string) *SetRenewalRequest {
  s.RenewalStatus = &v
  return s
}

type SetRenewalResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
}

func (s SetRenewalResponse) String() string {
  return tea.Prettify(s)
}

func (s SetRenewalResponse) GoString() string {
  return s.String()
}

func (s *SetRenewalResponse) SetRequestId(v string) *SetRenewalResponse {
  s.RequestId = &v
  return s
}

func (s *SetRenewalResponse) SetSuccess(v bool) *SetRenewalResponse {
  s.Success = &v
  return s
}

func (s *SetRenewalResponse) SetCode(v string) *SetRenewalResponse {
  s.Code = &v
  return s
}

func (s *SetRenewalResponse) SetMessage(v string) *SetRenewalResponse {
  s.Message = &v
  return s
}

type QueryAvailableInstancesRequest struct {
  Region *string `json:"Region" xml:"Region"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  PageNum *int `json:"PageNum" xml:"PageNum"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType"`
  InstanceIDs *string `json:"InstanceIDs" xml:"InstanceIDs"`
  EndTimeStart *string `json:"EndTimeStart" xml:"EndTimeStart"`
  EndTimeEnd *string `json:"EndTimeEnd" xml:"EndTimeEnd"`
  CreateTimeStart *string `json:"CreateTimeStart" xml:"CreateTimeStart"`
  CreateTimeEnd *string `json:"CreateTimeEnd" xml:"CreateTimeEnd"`
  RenewStatus *string `json:"RenewStatus" xml:"RenewStatus"`
}

func (s QueryAvailableInstancesRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAvailableInstancesRequest) GoString() string {
  return s.String()
}

func (s *QueryAvailableInstancesRequest) SetRegion(v string) *QueryAvailableInstancesRequest {
  s.Region = &v
  return s
}

func (s *QueryAvailableInstancesRequest) SetOwnerId(v int64) *QueryAvailableInstancesRequest {
  s.OwnerId = &v
  return s
}

func (s *QueryAvailableInstancesRequest) SetPageNum(v int) *QueryAvailableInstancesRequest {
  s.PageNum = &v
  return s
}

func (s *QueryAvailableInstancesRequest) SetPageSize(v int) *QueryAvailableInstancesRequest {
  s.PageSize = &v
  return s
}

func (s *QueryAvailableInstancesRequest) SetProductCode(v string) *QueryAvailableInstancesRequest {
  s.ProductCode = &v
  return s
}

func (s *QueryAvailableInstancesRequest) SetProductType(v string) *QueryAvailableInstancesRequest {
  s.ProductType = &v
  return s
}

func (s *QueryAvailableInstancesRequest) SetSubscriptionType(v string) *QueryAvailableInstancesRequest {
  s.SubscriptionType = &v
  return s
}

func (s *QueryAvailableInstancesRequest) SetInstanceIDs(v string) *QueryAvailableInstancesRequest {
  s.InstanceIDs = &v
  return s
}

func (s *QueryAvailableInstancesRequest) SetEndTimeStart(v string) *QueryAvailableInstancesRequest {
  s.EndTimeStart = &v
  return s
}

func (s *QueryAvailableInstancesRequest) SetEndTimeEnd(v string) *QueryAvailableInstancesRequest {
  s.EndTimeEnd = &v
  return s
}

func (s *QueryAvailableInstancesRequest) SetCreateTimeStart(v string) *QueryAvailableInstancesRequest {
  s.CreateTimeStart = &v
  return s
}

func (s *QueryAvailableInstancesRequest) SetCreateTimeEnd(v string) *QueryAvailableInstancesRequest {
  s.CreateTimeEnd = &v
  return s
}

func (s *QueryAvailableInstancesRequest) SetRenewStatus(v string) *QueryAvailableInstancesRequest {
  s.RenewStatus = &v
  return s
}

type QueryAvailableInstancesResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryAvailableInstancesResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryAvailableInstancesResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAvailableInstancesResponse) GoString() string {
  return s.String()
}

func (s *QueryAvailableInstancesResponse) SetRequestId(v string) *QueryAvailableInstancesResponse {
  s.RequestId = &v
  return s
}

func (s *QueryAvailableInstancesResponse) SetSuccess(v bool) *QueryAvailableInstancesResponse {
  s.Success = &v
  return s
}

func (s *QueryAvailableInstancesResponse) SetCode(v string) *QueryAvailableInstancesResponse {
  s.Code = &v
  return s
}

func (s *QueryAvailableInstancesResponse) SetMessage(v string) *QueryAvailableInstancesResponse {
  s.Message = &v
  return s
}

func (s *QueryAvailableInstancesResponse) SetData(v *QueryAvailableInstancesResponseData) *QueryAvailableInstancesResponse {
  s.Data = v
  return s
}

type QueryAvailableInstancesResponseData struct {
  PageNum *int `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  TotalCount *int `json:"TotalCount" xml:"TotalCount" require:"true"`
  InstanceList []*QueryAvailableInstancesResponseDataInstanceList `json:"InstanceList" xml:"InstanceList" require:"true" type:"Repeated"`
}

func (s QueryAvailableInstancesResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryAvailableInstancesResponseData) GoString() string {
  return s.String()
}

func (s *QueryAvailableInstancesResponseData) SetPageNum(v int) *QueryAvailableInstancesResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryAvailableInstancesResponseData) SetPageSize(v int) *QueryAvailableInstancesResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryAvailableInstancesResponseData) SetTotalCount(v int) *QueryAvailableInstancesResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryAvailableInstancesResponseData) SetInstanceList(v []*QueryAvailableInstancesResponseDataInstanceList) *QueryAvailableInstancesResponseData {
  s.InstanceList = v
  return s
}

type QueryAvailableInstancesResponseDataInstanceList struct     {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId" require:"true"`
  SellerId *int64 `json:"SellerId" xml:"SellerId" require:"true"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType" require:"true"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  InstanceID *string `json:"InstanceID" xml:"InstanceID" require:"true"`
  Region *string `json:"Region" xml:"Region" require:"true"`
  CreateTime *string `json:"CreateTime" xml:"CreateTime" require:"true"`
  EndTime *string `json:"EndTime" xml:"EndTime" require:"true"`
  StopTime *string `json:"StopTime" xml:"StopTime" require:"true"`
  ReleaseTime *string `json:"ReleaseTime" xml:"ReleaseTime" require:"true"`
  ExpectedReleaseTime *string `json:"ExpectedReleaseTime" xml:"ExpectedReleaseTime" require:"true"`
  Status *string `json:"Status" xml:"Status" require:"true"`
  SubStatus *string `json:"SubStatus" xml:"SubStatus" require:"true"`
  RenewStatus *string `json:"RenewStatus" xml:"RenewStatus" require:"true"`
  RenewalDuration *int `json:"RenewalDuration" xml:"RenewalDuration" require:"true"`
  RenewalDurationUnit *string `json:"RenewalDurationUnit" xml:"RenewalDurationUnit" require:"true"`
  Seller *string `json:"Seller" xml:"Seller" require:"true"`
}

func (s QueryAvailableInstancesResponseDataInstanceList) String() string {
  return tea.Prettify(s)
}

func (s QueryAvailableInstancesResponseDataInstanceList) GoString() string {
  return s.String()
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetOwnerId(v int64) *QueryAvailableInstancesResponseDataInstanceList {
  s.OwnerId = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetSellerId(v int64) *QueryAvailableInstancesResponseDataInstanceList {
  s.SellerId = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetProductCode(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.ProductCode = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetProductType(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.ProductType = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetSubscriptionType(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.SubscriptionType = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetInstanceID(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.InstanceID = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetRegion(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.Region = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetCreateTime(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.CreateTime = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetEndTime(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.EndTime = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetStopTime(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.StopTime = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetReleaseTime(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.ReleaseTime = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetExpectedReleaseTime(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.ExpectedReleaseTime = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetStatus(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.Status = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetSubStatus(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.SubStatus = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetRenewStatus(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.RenewStatus = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetRenewalDuration(v int) *QueryAvailableInstancesResponseDataInstanceList {
  s.RenewalDuration = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetRenewalDurationUnit(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.RenewalDurationUnit = &v
  return s
}

func (s *QueryAvailableInstancesResponseDataInstanceList) SetSeller(v string) *QueryAvailableInstancesResponseDataInstanceList {
  s.Seller = &v
  return s
}

type CreateResourcePackageRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode"`
  PackageType *string `json:"PackageType" xml:"PackageType"`
  EffectiveDate *string `json:"EffectiveDate" xml:"EffectiveDate"`
  Specification *string `json:"Specification" xml:"Specification"`
  Duration *int `json:"Duration" xml:"Duration"`
  PricingCycle *string `json:"PricingCycle" xml:"PricingCycle"`
}

func (s CreateResourcePackageRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateResourcePackageRequest) GoString() string {
  return s.String()
}

func (s *CreateResourcePackageRequest) SetOwnerId(v int64) *CreateResourcePackageRequest {
  s.OwnerId = &v
  return s
}

func (s *CreateResourcePackageRequest) SetProductCode(v string) *CreateResourcePackageRequest {
  s.ProductCode = &v
  return s
}

func (s *CreateResourcePackageRequest) SetPackageType(v string) *CreateResourcePackageRequest {
  s.PackageType = &v
  return s
}

func (s *CreateResourcePackageRequest) SetEffectiveDate(v string) *CreateResourcePackageRequest {
  s.EffectiveDate = &v
  return s
}

func (s *CreateResourcePackageRequest) SetSpecification(v string) *CreateResourcePackageRequest {
  s.Specification = &v
  return s
}

func (s *CreateResourcePackageRequest) SetDuration(v int) *CreateResourcePackageRequest {
  s.Duration = &v
  return s
}

func (s *CreateResourcePackageRequest) SetPricingCycle(v string) *CreateResourcePackageRequest {
  s.PricingCycle = &v
  return s
}

type CreateResourcePackageResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  OrderId *int64 `json:"OrderId" xml:"OrderId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *CreateResourcePackageResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s CreateResourcePackageResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateResourcePackageResponse) GoString() string {
  return s.String()
}

func (s *CreateResourcePackageResponse) SetRequestId(v string) *CreateResourcePackageResponse {
  s.RequestId = &v
  return s
}

func (s *CreateResourcePackageResponse) SetOrderId(v int64) *CreateResourcePackageResponse {
  s.OrderId = &v
  return s
}

func (s *CreateResourcePackageResponse) SetSuccess(v bool) *CreateResourcePackageResponse {
  s.Success = &v
  return s
}

func (s *CreateResourcePackageResponse) SetCode(v string) *CreateResourcePackageResponse {
  s.Code = &v
  return s
}

func (s *CreateResourcePackageResponse) SetMessage(v string) *CreateResourcePackageResponse {
  s.Message = &v
  return s
}

func (s *CreateResourcePackageResponse) SetData(v *CreateResourcePackageResponseData) *CreateResourcePackageResponse {
  s.Data = v
  return s
}

type CreateResourcePackageResponseData struct {
  OrderId *int64 `json:"OrderId" xml:"OrderId" require:"true"`
  InstanceId *string `json:"InstanceId" xml:"InstanceId" require:"true"`
}

func (s CreateResourcePackageResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateResourcePackageResponseData) GoString() string {
  return s.String()
}

func (s *CreateResourcePackageResponseData) SetOrderId(v int64) *CreateResourcePackageResponseData {
  s.OrderId = &v
  return s
}

func (s *CreateResourcePackageResponseData) SetInstanceId(v string) *CreateResourcePackageResponseData {
  s.InstanceId = &v
  return s
}

type QueryResourcePackageInstancesRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode"`
  ExpiryTimeStart *string `json:"ExpiryTimeStart" xml:"ExpiryTimeStart"`
  ExpiryTimeEnd *string `json:"ExpiryTimeEnd" xml:"ExpiryTimeEnd"`
  PageNum *int `json:"PageNum" xml:"PageNum"`
  PageSize *int `json:"PageSize" xml:"PageSize"`
}

func (s QueryResourcePackageInstancesRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryResourcePackageInstancesRequest) GoString() string {
  return s.String()
}

func (s *QueryResourcePackageInstancesRequest) SetOwnerId(v int64) *QueryResourcePackageInstancesRequest {
  s.OwnerId = &v
  return s
}

func (s *QueryResourcePackageInstancesRequest) SetProductCode(v string) *QueryResourcePackageInstancesRequest {
  s.ProductCode = &v
  return s
}

func (s *QueryResourcePackageInstancesRequest) SetExpiryTimeStart(v string) *QueryResourcePackageInstancesRequest {
  s.ExpiryTimeStart = &v
  return s
}

func (s *QueryResourcePackageInstancesRequest) SetExpiryTimeEnd(v string) *QueryResourcePackageInstancesRequest {
  s.ExpiryTimeEnd = &v
  return s
}

func (s *QueryResourcePackageInstancesRequest) SetPageNum(v int) *QueryResourcePackageInstancesRequest {
  s.PageNum = &v
  return s
}

func (s *QueryResourcePackageInstancesRequest) SetPageSize(v int) *QueryResourcePackageInstancesRequest {
  s.PageSize = &v
  return s
}

type QueryResourcePackageInstancesResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Page *int `json:"Page" xml:"Page" require:"true"`
  PageSize *int `json:"PageSize" xml:"PageSize" require:"true"`
  Total *int `json:"Total" xml:"Total" require:"true"`
  Data *QueryResourcePackageInstancesResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryResourcePackageInstancesResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryResourcePackageInstancesResponse) GoString() string {
  return s.String()
}

func (s *QueryResourcePackageInstancesResponse) SetRequestId(v string) *QueryResourcePackageInstancesResponse {
  s.RequestId = &v
  return s
}

func (s *QueryResourcePackageInstancesResponse) SetSuccess(v bool) *QueryResourcePackageInstancesResponse {
  s.Success = &v
  return s
}

func (s *QueryResourcePackageInstancesResponse) SetCode(v string) *QueryResourcePackageInstancesResponse {
  s.Code = &v
  return s
}

func (s *QueryResourcePackageInstancesResponse) SetMessage(v string) *QueryResourcePackageInstancesResponse {
  s.Message = &v
  return s
}

func (s *QueryResourcePackageInstancesResponse) SetPage(v int) *QueryResourcePackageInstancesResponse {
  s.Page = &v
  return s
}

func (s *QueryResourcePackageInstancesResponse) SetPageSize(v int) *QueryResourcePackageInstancesResponse {
  s.PageSize = &v
  return s
}

func (s *QueryResourcePackageInstancesResponse) SetTotal(v int) *QueryResourcePackageInstancesResponse {
  s.Total = &v
  return s
}

func (s *QueryResourcePackageInstancesResponse) SetData(v *QueryResourcePackageInstancesResponseData) *QueryResourcePackageInstancesResponse {
  s.Data = v
  return s
}

type QueryResourcePackageInstancesResponseData struct {
  HostId *string `json:"HostId" xml:"HostId" require:"true"`
  PageNum *string `json:"PageNum" xml:"PageNum" require:"true"`
  PageSize *string `json:"PageSize" xml:"PageSize" require:"true"`
  TotalCount *string `json:"TotalCount" xml:"TotalCount" require:"true"`
  Instances *QueryResourcePackageInstancesResponseDataInstances `json:"Instances" xml:"Instances" require:"true" type:"Struct"`
}

func (s QueryResourcePackageInstancesResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryResourcePackageInstancesResponseData) GoString() string {
  return s.String()
}

func (s *QueryResourcePackageInstancesResponseData) SetHostId(v string) *QueryResourcePackageInstancesResponseData {
  s.HostId = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseData) SetPageNum(v string) *QueryResourcePackageInstancesResponseData {
  s.PageNum = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseData) SetPageSize(v string) *QueryResourcePackageInstancesResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseData) SetTotalCount(v string) *QueryResourcePackageInstancesResponseData {
  s.TotalCount = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseData) SetInstances(v *QueryResourcePackageInstancesResponseDataInstances) *QueryResourcePackageInstancesResponseData {
  s.Instances = v
  return s
}

type QueryResourcePackageInstancesResponseDataInstances struct {
  Instance []*QueryResourcePackageInstancesResponseDataInstancesInstance `json:"Instance" xml:"Instance" require:"true" type:"Repeated"`
}

func (s QueryResourcePackageInstancesResponseDataInstances) String() string {
  return tea.Prettify(s)
}

func (s QueryResourcePackageInstancesResponseDataInstances) GoString() string {
  return s.String()
}

func (s *QueryResourcePackageInstancesResponseDataInstances) SetInstance(v []*QueryResourcePackageInstancesResponseDataInstancesInstance) *QueryResourcePackageInstancesResponseDataInstances {
  s.Instance = v
  return s
}

type QueryResourcePackageInstancesResponseDataInstancesInstance struct     {
  InstanceId *string `json:"InstanceId" xml:"InstanceId" require:"true"`
  Region *string `json:"Region" xml:"Region" require:"true"`
  TotalAmount *string `json:"TotalAmount" xml:"TotalAmount" require:"true"`
  TotalAmountUnit *string `json:"TotalAmountUnit" xml:"TotalAmountUnit" require:"true"`
  RemainingAmount *string `json:"RemainingAmount" xml:"RemainingAmount" require:"true"`
  RemainingAmountUnit *string `json:"RemainingAmountUnit" xml:"RemainingAmountUnit" require:"true"`
  EffectiveTime *string `json:"EffectiveTime" xml:"EffectiveTime" require:"true"`
  ExpiryTime *string `json:"ExpiryTime" xml:"ExpiryTime" require:"true"`
  Remark *string `json:"Remark" xml:"Remark" require:"true"`
  PackageType *string `json:"PackageType" xml:"PackageType" require:"true"`
  Status *string `json:"Status" xml:"Status" require:"true"`
  DeductType *string `json:"DeductType" xml:"DeductType" require:"true"`
  ApplicableProducts *QueryResourcePackageInstancesResponseDataInstancesInstanceApplicableProducts `json:"ApplicableProducts" xml:"ApplicableProducts" require:"true" type:"Struct"`
}

func (s QueryResourcePackageInstancesResponseDataInstancesInstance) String() string {
  return tea.Prettify(s)
}

func (s QueryResourcePackageInstancesResponseDataInstancesInstance) GoString() string {
  return s.String()
}

func (s *QueryResourcePackageInstancesResponseDataInstancesInstance) SetInstanceId(v string) *QueryResourcePackageInstancesResponseDataInstancesInstance {
  s.InstanceId = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseDataInstancesInstance) SetRegion(v string) *QueryResourcePackageInstancesResponseDataInstancesInstance {
  s.Region = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseDataInstancesInstance) SetTotalAmount(v string) *QueryResourcePackageInstancesResponseDataInstancesInstance {
  s.TotalAmount = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseDataInstancesInstance) SetTotalAmountUnit(v string) *QueryResourcePackageInstancesResponseDataInstancesInstance {
  s.TotalAmountUnit = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseDataInstancesInstance) SetRemainingAmount(v string) *QueryResourcePackageInstancesResponseDataInstancesInstance {
  s.RemainingAmount = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseDataInstancesInstance) SetRemainingAmountUnit(v string) *QueryResourcePackageInstancesResponseDataInstancesInstance {
  s.RemainingAmountUnit = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseDataInstancesInstance) SetEffectiveTime(v string) *QueryResourcePackageInstancesResponseDataInstancesInstance {
  s.EffectiveTime = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseDataInstancesInstance) SetExpiryTime(v string) *QueryResourcePackageInstancesResponseDataInstancesInstance {
  s.ExpiryTime = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseDataInstancesInstance) SetRemark(v string) *QueryResourcePackageInstancesResponseDataInstancesInstance {
  s.Remark = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseDataInstancesInstance) SetPackageType(v string) *QueryResourcePackageInstancesResponseDataInstancesInstance {
  s.PackageType = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseDataInstancesInstance) SetStatus(v string) *QueryResourcePackageInstancesResponseDataInstancesInstance {
  s.Status = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseDataInstancesInstance) SetDeductType(v string) *QueryResourcePackageInstancesResponseDataInstancesInstance {
  s.DeductType = &v
  return s
}

func (s *QueryResourcePackageInstancesResponseDataInstancesInstance) SetApplicableProducts(v *QueryResourcePackageInstancesResponseDataInstancesInstanceApplicableProducts) *QueryResourcePackageInstancesResponseDataInstancesInstance {
  s.ApplicableProducts = v
  return s
}

type QueryResourcePackageInstancesResponseDataInstancesInstanceApplicableProducts struct {
  Product []*string `json:"Product" xml:"Product" require:"true" type:"Repeated"`
}

func (s QueryResourcePackageInstancesResponseDataInstancesInstanceApplicableProducts) String() string {
  return tea.Prettify(s)
}

func (s QueryResourcePackageInstancesResponseDataInstancesInstanceApplicableProducts) GoString() string {
  return s.String()
}

func (s *QueryResourcePackageInstancesResponseDataInstancesInstanceApplicableProducts) SetProduct(v []*string) *QueryResourcePackageInstancesResponseDataInstancesInstanceApplicableProducts {
  s.Product = v
  return s
}

type GetResourcePackagePriceRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode"`
  PackageType *string `json:"PackageType" xml:"PackageType"`
  EffectiveDate *string `json:"EffectiveDate" xml:"EffectiveDate"`
  Specification *string `json:"Specification" xml:"Specification"`
  Duration *int `json:"Duration" xml:"Duration"`
  PricingCycle *string `json:"PricingCycle" xml:"PricingCycle"`
  OrderType *string `json:"OrderType" xml:"OrderType"`
  InstanceId *string `json:"InstanceId" xml:"InstanceId"`
}

func (s GetResourcePackagePriceRequest) String() string {
  return tea.Prettify(s)
}

func (s GetResourcePackagePriceRequest) GoString() string {
  return s.String()
}

func (s *GetResourcePackagePriceRequest) SetOwnerId(v int64) *GetResourcePackagePriceRequest {
  s.OwnerId = &v
  return s
}

func (s *GetResourcePackagePriceRequest) SetProductCode(v string) *GetResourcePackagePriceRequest {
  s.ProductCode = &v
  return s
}

func (s *GetResourcePackagePriceRequest) SetPackageType(v string) *GetResourcePackagePriceRequest {
  s.PackageType = &v
  return s
}

func (s *GetResourcePackagePriceRequest) SetEffectiveDate(v string) *GetResourcePackagePriceRequest {
  s.EffectiveDate = &v
  return s
}

func (s *GetResourcePackagePriceRequest) SetSpecification(v string) *GetResourcePackagePriceRequest {
  s.Specification = &v
  return s
}

func (s *GetResourcePackagePriceRequest) SetDuration(v int) *GetResourcePackagePriceRequest {
  s.Duration = &v
  return s
}

func (s *GetResourcePackagePriceRequest) SetPricingCycle(v string) *GetResourcePackagePriceRequest {
  s.PricingCycle = &v
  return s
}

func (s *GetResourcePackagePriceRequest) SetOrderType(v string) *GetResourcePackagePriceRequest {
  s.OrderType = &v
  return s
}

func (s *GetResourcePackagePriceRequest) SetInstanceId(v string) *GetResourcePackagePriceRequest {
  s.InstanceId = &v
  return s
}

type GetResourcePackagePriceResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *GetResourcePackagePriceResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s GetResourcePackagePriceResponse) String() string {
  return tea.Prettify(s)
}

func (s GetResourcePackagePriceResponse) GoString() string {
  return s.String()
}

func (s *GetResourcePackagePriceResponse) SetRequestId(v string) *GetResourcePackagePriceResponse {
  s.RequestId = &v
  return s
}

func (s *GetResourcePackagePriceResponse) SetSuccess(v bool) *GetResourcePackagePriceResponse {
  s.Success = &v
  return s
}

func (s *GetResourcePackagePriceResponse) SetCode(v string) *GetResourcePackagePriceResponse {
  s.Code = &v
  return s
}

func (s *GetResourcePackagePriceResponse) SetMessage(v string) *GetResourcePackagePriceResponse {
  s.Message = &v
  return s
}

func (s *GetResourcePackagePriceResponse) SetData(v *GetResourcePackagePriceResponseData) *GetResourcePackagePriceResponse {
  s.Data = v
  return s
}

type GetResourcePackagePriceResponseData struct {
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  OriginalPrice *float32 `json:"OriginalPrice" xml:"OriginalPrice" require:"true"`
  TradePrice *float32 `json:"TradePrice" xml:"TradePrice" require:"true"`
  DiscountPrice *float32 `json:"DiscountPrice" xml:"DiscountPrice" require:"true"`
  Promotions *GetResourcePackagePriceResponseDataPromotions `json:"Promotions" xml:"Promotions" require:"true" type:"Struct"`
}

func (s GetResourcePackagePriceResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetResourcePackagePriceResponseData) GoString() string {
  return s.String()
}

func (s *GetResourcePackagePriceResponseData) SetCurrency(v string) *GetResourcePackagePriceResponseData {
  s.Currency = &v
  return s
}

func (s *GetResourcePackagePriceResponseData) SetOriginalPrice(v float32) *GetResourcePackagePriceResponseData {
  s.OriginalPrice = &v
  return s
}

func (s *GetResourcePackagePriceResponseData) SetTradePrice(v float32) *GetResourcePackagePriceResponseData {
  s.TradePrice = &v
  return s
}

func (s *GetResourcePackagePriceResponseData) SetDiscountPrice(v float32) *GetResourcePackagePriceResponseData {
  s.DiscountPrice = &v
  return s
}

func (s *GetResourcePackagePriceResponseData) SetPromotions(v *GetResourcePackagePriceResponseDataPromotions) *GetResourcePackagePriceResponseData {
  s.Promotions = v
  return s
}

type GetResourcePackagePriceResponseDataPromotions struct {
  Promotion []*GetResourcePackagePriceResponseDataPromotionsPromotion `json:"Promotion" xml:"Promotion" require:"true" type:"Repeated"`
}

func (s GetResourcePackagePriceResponseDataPromotions) String() string {
  return tea.Prettify(s)
}

func (s GetResourcePackagePriceResponseDataPromotions) GoString() string {
  return s.String()
}

func (s *GetResourcePackagePriceResponseDataPromotions) SetPromotion(v []*GetResourcePackagePriceResponseDataPromotionsPromotion) *GetResourcePackagePriceResponseDataPromotions {
  s.Promotion = v
  return s
}

type GetResourcePackagePriceResponseDataPromotionsPromotion struct     {
  Id *int64 `json:"Id" xml:"Id" require:"true"`
  Name *string `json:"Name" xml:"Name" require:"true"`
}

func (s GetResourcePackagePriceResponseDataPromotionsPromotion) String() string {
  return tea.Prettify(s)
}

func (s GetResourcePackagePriceResponseDataPromotionsPromotion) GoString() string {
  return s.String()
}

func (s *GetResourcePackagePriceResponseDataPromotionsPromotion) SetId(v int64) *GetResourcePackagePriceResponseDataPromotionsPromotion {
  s.Id = &v
  return s
}

func (s *GetResourcePackagePriceResponseDataPromotionsPromotion) SetName(v string) *GetResourcePackagePriceResponseDataPromotionsPromotion {
  s.Name = &v
  return s
}

type GetSubscriptionPriceRequest struct {
  ServicePeriodUnit *string `json:"ServicePeriodUnit" xml:"ServicePeriodUnit"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  OrderType *string `json:"OrderType" xml:"OrderType" require:"true"`
  ServicePeriodQuantity *int `json:"ServicePeriodQuantity" xml:"ServicePeriodQuantity"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  Region *string `json:"Region" xml:"Region"`
  InstanceId *string `json:"InstanceId" xml:"InstanceId"`
  ModuleList []*GetSubscriptionPriceRequestModuleList `json:"ModuleList" xml:"ModuleList" require:"true" type:"Repeated"`
  Quantity *int `json:"Quantity" xml:"Quantity"`
}

func (s GetSubscriptionPriceRequest) String() string {
  return tea.Prettify(s)
}

func (s GetSubscriptionPriceRequest) GoString() string {
  return s.String()
}

func (s *GetSubscriptionPriceRequest) SetServicePeriodUnit(v string) *GetSubscriptionPriceRequest {
  s.ServicePeriodUnit = &v
  return s
}

func (s *GetSubscriptionPriceRequest) SetSubscriptionType(v string) *GetSubscriptionPriceRequest {
  s.SubscriptionType = &v
  return s
}

func (s *GetSubscriptionPriceRequest) SetOwnerId(v int64) *GetSubscriptionPriceRequest {
  s.OwnerId = &v
  return s
}

func (s *GetSubscriptionPriceRequest) SetProductCode(v string) *GetSubscriptionPriceRequest {
  s.ProductCode = &v
  return s
}

func (s *GetSubscriptionPriceRequest) SetOrderType(v string) *GetSubscriptionPriceRequest {
  s.OrderType = &v
  return s
}

func (s *GetSubscriptionPriceRequest) SetServicePeriodQuantity(v int) *GetSubscriptionPriceRequest {
  s.ServicePeriodQuantity = &v
  return s
}

func (s *GetSubscriptionPriceRequest) SetProductType(v string) *GetSubscriptionPriceRequest {
  s.ProductType = &v
  return s
}

func (s *GetSubscriptionPriceRequest) SetRegion(v string) *GetSubscriptionPriceRequest {
  s.Region = &v
  return s
}

func (s *GetSubscriptionPriceRequest) SetInstanceId(v string) *GetSubscriptionPriceRequest {
  s.InstanceId = &v
  return s
}

func (s *GetSubscriptionPriceRequest) SetModuleList(v []*GetSubscriptionPriceRequestModuleList) *GetSubscriptionPriceRequest {
  s.ModuleList = v
  return s
}

func (s *GetSubscriptionPriceRequest) SetQuantity(v int) *GetSubscriptionPriceRequest {
  s.Quantity = &v
  return s
}

type GetSubscriptionPriceRequestModuleList struct     {
  ModuleCode *string `json:"ModuleCode" xml:"ModuleCode" require:"true"`
  Config *string `json:"Config" xml:"Config" require:"true"`
  ModuleStatus *int `json:"ModuleStatus" xml:"ModuleStatus"`
  Tag *string `json:"Tag" xml:"Tag"`
}

func (s GetSubscriptionPriceRequestModuleList) String() string {
  return tea.Prettify(s)
}

func (s GetSubscriptionPriceRequestModuleList) GoString() string {
  return s.String()
}

func (s *GetSubscriptionPriceRequestModuleList) SetModuleCode(v string) *GetSubscriptionPriceRequestModuleList {
  s.ModuleCode = &v
  return s
}

func (s *GetSubscriptionPriceRequestModuleList) SetConfig(v string) *GetSubscriptionPriceRequestModuleList {
  s.Config = &v
  return s
}

func (s *GetSubscriptionPriceRequestModuleList) SetModuleStatus(v int) *GetSubscriptionPriceRequestModuleList {
  s.ModuleStatus = &v
  return s
}

func (s *GetSubscriptionPriceRequestModuleList) SetTag(v string) *GetSubscriptionPriceRequestModuleList {
  s.Tag = &v
  return s
}

type GetSubscriptionPriceResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *GetSubscriptionPriceResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s GetSubscriptionPriceResponse) String() string {
  return tea.Prettify(s)
}

func (s GetSubscriptionPriceResponse) GoString() string {
  return s.String()
}

func (s *GetSubscriptionPriceResponse) SetRequestId(v string) *GetSubscriptionPriceResponse {
  s.RequestId = &v
  return s
}

func (s *GetSubscriptionPriceResponse) SetSuccess(v bool) *GetSubscriptionPriceResponse {
  s.Success = &v
  return s
}

func (s *GetSubscriptionPriceResponse) SetCode(v string) *GetSubscriptionPriceResponse {
  s.Code = &v
  return s
}

func (s *GetSubscriptionPriceResponse) SetMessage(v string) *GetSubscriptionPriceResponse {
  s.Message = &v
  return s
}

func (s *GetSubscriptionPriceResponse) SetData(v *GetSubscriptionPriceResponseData) *GetSubscriptionPriceResponse {
  s.Data = v
  return s
}

type GetSubscriptionPriceResponseData struct {
  OriginalPrice *float32 `json:"OriginalPrice" xml:"OriginalPrice" require:"true"`
  DiscountPrice *float32 `json:"DiscountPrice" xml:"DiscountPrice" require:"true"`
  TradePrice *float32 `json:"TradePrice" xml:"TradePrice" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  Quantity *int `json:"Quantity" xml:"Quantity" require:"true"`
  ModuleDetails *GetSubscriptionPriceResponseDataModuleDetails `json:"ModuleDetails" xml:"ModuleDetails" require:"true" type:"Struct"`
  PromotionDetails *GetSubscriptionPriceResponseDataPromotionDetails `json:"PromotionDetails" xml:"PromotionDetails" require:"true" type:"Struct"`
}

func (s GetSubscriptionPriceResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetSubscriptionPriceResponseData) GoString() string {
  return s.String()
}

func (s *GetSubscriptionPriceResponseData) SetOriginalPrice(v float32) *GetSubscriptionPriceResponseData {
  s.OriginalPrice = &v
  return s
}

func (s *GetSubscriptionPriceResponseData) SetDiscountPrice(v float32) *GetSubscriptionPriceResponseData {
  s.DiscountPrice = &v
  return s
}

func (s *GetSubscriptionPriceResponseData) SetTradePrice(v float32) *GetSubscriptionPriceResponseData {
  s.TradePrice = &v
  return s
}

func (s *GetSubscriptionPriceResponseData) SetCurrency(v string) *GetSubscriptionPriceResponseData {
  s.Currency = &v
  return s
}

func (s *GetSubscriptionPriceResponseData) SetQuantity(v int) *GetSubscriptionPriceResponseData {
  s.Quantity = &v
  return s
}

func (s *GetSubscriptionPriceResponseData) SetModuleDetails(v *GetSubscriptionPriceResponseDataModuleDetails) *GetSubscriptionPriceResponseData {
  s.ModuleDetails = v
  return s
}

func (s *GetSubscriptionPriceResponseData) SetPromotionDetails(v *GetSubscriptionPriceResponseDataPromotionDetails) *GetSubscriptionPriceResponseData {
  s.PromotionDetails = v
  return s
}

type GetSubscriptionPriceResponseDataModuleDetails struct {
  ModuleDetail []*GetSubscriptionPriceResponseDataModuleDetailsModuleDetail `json:"ModuleDetail" xml:"ModuleDetail" require:"true" type:"Repeated"`
}

func (s GetSubscriptionPriceResponseDataModuleDetails) String() string {
  return tea.Prettify(s)
}

func (s GetSubscriptionPriceResponseDataModuleDetails) GoString() string {
  return s.String()
}

func (s *GetSubscriptionPriceResponseDataModuleDetails) SetModuleDetail(v []*GetSubscriptionPriceResponseDataModuleDetailsModuleDetail) *GetSubscriptionPriceResponseDataModuleDetails {
  s.ModuleDetail = v
  return s
}

type GetSubscriptionPriceResponseDataModuleDetailsModuleDetail struct     {
  ModuleCode *string `json:"ModuleCode" xml:"ModuleCode" require:"true"`
  OriginalCost *float32 `json:"OriginalCost" xml:"OriginalCost" require:"true"`
  InvoiceDiscount *float32 `json:"InvoiceDiscount" xml:"InvoiceDiscount" require:"true"`
  CostAfterDiscount *float32 `json:"CostAfterDiscount" xml:"CostAfterDiscount" require:"true"`
  UnitPrice *float32 `json:"UnitPrice" xml:"UnitPrice" require:"true"`
}

func (s GetSubscriptionPriceResponseDataModuleDetailsModuleDetail) String() string {
  return tea.Prettify(s)
}

func (s GetSubscriptionPriceResponseDataModuleDetailsModuleDetail) GoString() string {
  return s.String()
}

func (s *GetSubscriptionPriceResponseDataModuleDetailsModuleDetail) SetModuleCode(v string) *GetSubscriptionPriceResponseDataModuleDetailsModuleDetail {
  s.ModuleCode = &v
  return s
}

func (s *GetSubscriptionPriceResponseDataModuleDetailsModuleDetail) SetOriginalCost(v float32) *GetSubscriptionPriceResponseDataModuleDetailsModuleDetail {
  s.OriginalCost = &v
  return s
}

func (s *GetSubscriptionPriceResponseDataModuleDetailsModuleDetail) SetInvoiceDiscount(v float32) *GetSubscriptionPriceResponseDataModuleDetailsModuleDetail {
  s.InvoiceDiscount = &v
  return s
}

func (s *GetSubscriptionPriceResponseDataModuleDetailsModuleDetail) SetCostAfterDiscount(v float32) *GetSubscriptionPriceResponseDataModuleDetailsModuleDetail {
  s.CostAfterDiscount = &v
  return s
}

func (s *GetSubscriptionPriceResponseDataModuleDetailsModuleDetail) SetUnitPrice(v float32) *GetSubscriptionPriceResponseDataModuleDetailsModuleDetail {
  s.UnitPrice = &v
  return s
}

type GetSubscriptionPriceResponseDataPromotionDetails struct {
  PromotionDetail []*GetSubscriptionPriceResponseDataPromotionDetailsPromotionDetail `json:"PromotionDetail" xml:"PromotionDetail" require:"true" type:"Repeated"`
}

func (s GetSubscriptionPriceResponseDataPromotionDetails) String() string {
  return tea.Prettify(s)
}

func (s GetSubscriptionPriceResponseDataPromotionDetails) GoString() string {
  return s.String()
}

func (s *GetSubscriptionPriceResponseDataPromotionDetails) SetPromotionDetail(v []*GetSubscriptionPriceResponseDataPromotionDetailsPromotionDetail) *GetSubscriptionPriceResponseDataPromotionDetails {
  s.PromotionDetail = v
  return s
}

type GetSubscriptionPriceResponseDataPromotionDetailsPromotionDetail struct     {
  PromotionName *string `json:"PromotionName" xml:"PromotionName" require:"true"`
  PromotionDesc *string `json:"PromotionDesc" xml:"PromotionDesc" require:"true"`
  PromotionId *int64 `json:"PromotionId" xml:"PromotionId" require:"true"`
}

func (s GetSubscriptionPriceResponseDataPromotionDetailsPromotionDetail) String() string {
  return tea.Prettify(s)
}

func (s GetSubscriptionPriceResponseDataPromotionDetailsPromotionDetail) GoString() string {
  return s.String()
}

func (s *GetSubscriptionPriceResponseDataPromotionDetailsPromotionDetail) SetPromotionName(v string) *GetSubscriptionPriceResponseDataPromotionDetailsPromotionDetail {
  s.PromotionName = &v
  return s
}

func (s *GetSubscriptionPriceResponseDataPromotionDetailsPromotionDetail) SetPromotionDesc(v string) *GetSubscriptionPriceResponseDataPromotionDetailsPromotionDetail {
  s.PromotionDesc = &v
  return s
}

func (s *GetSubscriptionPriceResponseDataPromotionDetailsPromotionDetail) SetPromotionId(v int64) *GetSubscriptionPriceResponseDataPromotionDetailsPromotionDetail {
  s.PromotionId = &v
  return s
}

type GetPayAsYouGoPriceRequest struct {
  OwnerId *int64 `json:"OwnerId" xml:"OwnerId"`
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType"`
  SubscriptionType *string `json:"SubscriptionType" xml:"SubscriptionType" require:"true"`
  Region *string `json:"Region" xml:"Region"`
  ModuleList []*GetPayAsYouGoPriceRequestModuleList `json:"ModuleList" xml:"ModuleList" require:"true" type:"Repeated"`
}

func (s GetPayAsYouGoPriceRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceRequest) GoString() string {
  return s.String()
}

func (s *GetPayAsYouGoPriceRequest) SetOwnerId(v int64) *GetPayAsYouGoPriceRequest {
  s.OwnerId = &v
  return s
}

func (s *GetPayAsYouGoPriceRequest) SetProductCode(v string) *GetPayAsYouGoPriceRequest {
  s.ProductCode = &v
  return s
}

func (s *GetPayAsYouGoPriceRequest) SetProductType(v string) *GetPayAsYouGoPriceRequest {
  s.ProductType = &v
  return s
}

func (s *GetPayAsYouGoPriceRequest) SetSubscriptionType(v string) *GetPayAsYouGoPriceRequest {
  s.SubscriptionType = &v
  return s
}

func (s *GetPayAsYouGoPriceRequest) SetRegion(v string) *GetPayAsYouGoPriceRequest {
  s.Region = &v
  return s
}

func (s *GetPayAsYouGoPriceRequest) SetModuleList(v []*GetPayAsYouGoPriceRequestModuleList) *GetPayAsYouGoPriceRequest {
  s.ModuleList = v
  return s
}

type GetPayAsYouGoPriceRequestModuleList struct     {
  ModuleCode *string `json:"ModuleCode" xml:"ModuleCode" require:"true"`
  Config *string `json:"Config" xml:"Config" require:"true"`
  PriceType *string `json:"PriceType" xml:"PriceType" require:"true"`
}

func (s GetPayAsYouGoPriceRequestModuleList) String() string {
  return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceRequestModuleList) GoString() string {
  return s.String()
}

func (s *GetPayAsYouGoPriceRequestModuleList) SetModuleCode(v string) *GetPayAsYouGoPriceRequestModuleList {
  s.ModuleCode = &v
  return s
}

func (s *GetPayAsYouGoPriceRequestModuleList) SetConfig(v string) *GetPayAsYouGoPriceRequestModuleList {
  s.Config = &v
  return s
}

func (s *GetPayAsYouGoPriceRequestModuleList) SetPriceType(v string) *GetPayAsYouGoPriceRequestModuleList {
  s.PriceType = &v
  return s
}

type GetPayAsYouGoPriceResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *GetPayAsYouGoPriceResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s GetPayAsYouGoPriceResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceResponse) GoString() string {
  return s.String()
}

func (s *GetPayAsYouGoPriceResponse) SetRequestId(v string) *GetPayAsYouGoPriceResponse {
  s.RequestId = &v
  return s
}

func (s *GetPayAsYouGoPriceResponse) SetSuccess(v bool) *GetPayAsYouGoPriceResponse {
  s.Success = &v
  return s
}

func (s *GetPayAsYouGoPriceResponse) SetCode(v string) *GetPayAsYouGoPriceResponse {
  s.Code = &v
  return s
}

func (s *GetPayAsYouGoPriceResponse) SetMessage(v string) *GetPayAsYouGoPriceResponse {
  s.Message = &v
  return s
}

func (s *GetPayAsYouGoPriceResponse) SetData(v *GetPayAsYouGoPriceResponseData) *GetPayAsYouGoPriceResponse {
  s.Data = v
  return s
}

type GetPayAsYouGoPriceResponseData struct {
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
  ModuleDetails *GetPayAsYouGoPriceResponseDataModuleDetails `json:"ModuleDetails" xml:"ModuleDetails" require:"true" type:"Struct"`
  PromotionDetails *GetPayAsYouGoPriceResponseDataPromotionDetails `json:"PromotionDetails" xml:"PromotionDetails" require:"true" type:"Struct"`
}

func (s GetPayAsYouGoPriceResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseData) GoString() string {
  return s.String()
}

func (s *GetPayAsYouGoPriceResponseData) SetCurrency(v string) *GetPayAsYouGoPriceResponseData {
  s.Currency = &v
  return s
}

func (s *GetPayAsYouGoPriceResponseData) SetModuleDetails(v *GetPayAsYouGoPriceResponseDataModuleDetails) *GetPayAsYouGoPriceResponseData {
  s.ModuleDetails = v
  return s
}

func (s *GetPayAsYouGoPriceResponseData) SetPromotionDetails(v *GetPayAsYouGoPriceResponseDataPromotionDetails) *GetPayAsYouGoPriceResponseData {
  s.PromotionDetails = v
  return s
}

type GetPayAsYouGoPriceResponseDataModuleDetails struct {
  ModuleDetail []*GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail `json:"ModuleDetail" xml:"ModuleDetail" require:"true" type:"Repeated"`
}

func (s GetPayAsYouGoPriceResponseDataModuleDetails) String() string {
  return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseDataModuleDetails) GoString() string {
  return s.String()
}

func (s *GetPayAsYouGoPriceResponseDataModuleDetails) SetModuleDetail(v []*GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail) *GetPayAsYouGoPriceResponseDataModuleDetails {
  s.ModuleDetail = v
  return s
}

type GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail struct     {
  ModuleCode *string `json:"ModuleCode" xml:"ModuleCode" require:"true"`
  OriginalCost *float32 `json:"OriginalCost" xml:"OriginalCost" require:"true"`
  InvoiceDiscount *float32 `json:"InvoiceDiscount" xml:"InvoiceDiscount" require:"true"`
  CostAfterDiscount *float32 `json:"CostAfterDiscount" xml:"CostAfterDiscount" require:"true"`
  UnitPrice *float32 `json:"UnitPrice" xml:"UnitPrice" require:"true"`
}

func (s GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail) String() string {
  return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail) GoString() string {
  return s.String()
}

func (s *GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail) SetModuleCode(v string) *GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail {
  s.ModuleCode = &v
  return s
}

func (s *GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail) SetOriginalCost(v float32) *GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail {
  s.OriginalCost = &v
  return s
}

func (s *GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail) SetInvoiceDiscount(v float32) *GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail {
  s.InvoiceDiscount = &v
  return s
}

func (s *GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail) SetCostAfterDiscount(v float32) *GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail {
  s.CostAfterDiscount = &v
  return s
}

func (s *GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail) SetUnitPrice(v float32) *GetPayAsYouGoPriceResponseDataModuleDetailsModuleDetail {
  s.UnitPrice = &v
  return s
}

type GetPayAsYouGoPriceResponseDataPromotionDetails struct {
  PromotionDetail []*GetPayAsYouGoPriceResponseDataPromotionDetailsPromotionDetail `json:"PromotionDetail" xml:"PromotionDetail" require:"true" type:"Repeated"`
}

func (s GetPayAsYouGoPriceResponseDataPromotionDetails) String() string {
  return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseDataPromotionDetails) GoString() string {
  return s.String()
}

func (s *GetPayAsYouGoPriceResponseDataPromotionDetails) SetPromotionDetail(v []*GetPayAsYouGoPriceResponseDataPromotionDetailsPromotionDetail) *GetPayAsYouGoPriceResponseDataPromotionDetails {
  s.PromotionDetail = v
  return s
}

type GetPayAsYouGoPriceResponseDataPromotionDetailsPromotionDetail struct     {
  PromotionName *string `json:"PromotionName" xml:"PromotionName" require:"true"`
  PromotionDesc *string `json:"PromotionDesc" xml:"PromotionDesc" require:"true"`
  PromotionId *int64 `json:"PromotionId" xml:"PromotionId" require:"true"`
}

func (s GetPayAsYouGoPriceResponseDataPromotionDetailsPromotionDetail) String() string {
  return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseDataPromotionDetailsPromotionDetail) GoString() string {
  return s.String()
}

func (s *GetPayAsYouGoPriceResponseDataPromotionDetailsPromotionDetail) SetPromotionName(v string) *GetPayAsYouGoPriceResponseDataPromotionDetailsPromotionDetail {
  s.PromotionName = &v
  return s
}

func (s *GetPayAsYouGoPriceResponseDataPromotionDetailsPromotionDetail) SetPromotionDesc(v string) *GetPayAsYouGoPriceResponseDataPromotionDetailsPromotionDetail {
  s.PromotionDesc = &v
  return s
}

func (s *GetPayAsYouGoPriceResponseDataPromotionDetailsPromotionDetail) SetPromotionId(v int64) *GetPayAsYouGoPriceResponseDataPromotionDetailsPromotionDetail {
  s.PromotionId = &v
  return s
}

type QueryPrepaidCardsRequest struct {
  ExpiryTimeEnd *string `json:"ExpiryTimeEnd" xml:"ExpiryTimeEnd"`
  ExpiryTimeStart *string `json:"ExpiryTimeStart" xml:"ExpiryTimeStart"`
  EffectiveOrNot *bool `json:"EffectiveOrNot" xml:"EffectiveOrNot"`
}

func (s QueryPrepaidCardsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryPrepaidCardsRequest) GoString() string {
  return s.String()
}

func (s *QueryPrepaidCardsRequest) SetExpiryTimeEnd(v string) *QueryPrepaidCardsRequest {
  s.ExpiryTimeEnd = &v
  return s
}

func (s *QueryPrepaidCardsRequest) SetExpiryTimeStart(v string) *QueryPrepaidCardsRequest {
  s.ExpiryTimeStart = &v
  return s
}

func (s *QueryPrepaidCardsRequest) SetEffectiveOrNot(v bool) *QueryPrepaidCardsRequest {
  s.EffectiveOrNot = &v
  return s
}

type QueryPrepaidCardsResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryPrepaidCardsResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryPrepaidCardsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryPrepaidCardsResponse) GoString() string {
  return s.String()
}

func (s *QueryPrepaidCardsResponse) SetRequestId(v string) *QueryPrepaidCardsResponse {
  s.RequestId = &v
  return s
}

func (s *QueryPrepaidCardsResponse) SetSuccess(v bool) *QueryPrepaidCardsResponse {
  s.Success = &v
  return s
}

func (s *QueryPrepaidCardsResponse) SetCode(v string) *QueryPrepaidCardsResponse {
  s.Code = &v
  return s
}

func (s *QueryPrepaidCardsResponse) SetMessage(v string) *QueryPrepaidCardsResponse {
  s.Message = &v
  return s
}

func (s *QueryPrepaidCardsResponse) SetData(v *QueryPrepaidCardsResponseData) *QueryPrepaidCardsResponse {
  s.Data = v
  return s
}

type QueryPrepaidCardsResponseData struct {
  PrepaidCard []*QueryPrepaidCardsResponseDataPrepaidCard `json:"PrepaidCard" xml:"PrepaidCard" require:"true" type:"Repeated"`
}

func (s QueryPrepaidCardsResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryPrepaidCardsResponseData) GoString() string {
  return s.String()
}

func (s *QueryPrepaidCardsResponseData) SetPrepaidCard(v []*QueryPrepaidCardsResponseDataPrepaidCard) *QueryPrepaidCardsResponseData {
  s.PrepaidCard = v
  return s
}

type QueryPrepaidCardsResponseDataPrepaidCard struct     {
  PrepaidCardId *int64 `json:"PrepaidCardId" xml:"PrepaidCardId" require:"true"`
  PrepaidCardNo *string `json:"PrepaidCardNo" xml:"PrepaidCardNo" require:"true"`
  GrantedTime *string `json:"GrantedTime" xml:"GrantedTime" require:"true"`
  EffectiveTime *string `json:"EffectiveTime" xml:"EffectiveTime" require:"true"`
  ExpiryTime *string `json:"ExpiryTime" xml:"ExpiryTime" require:"true"`
  ApplicableProducts *string `json:"ApplicableProducts" xml:"ApplicableProducts" require:"true"`
  ApplicableScenarios *string `json:"ApplicableScenarios" xml:"ApplicableScenarios" require:"true"`
  NominalValue *string `json:"NominalValue" xml:"NominalValue" require:"true"`
  Balance *string `json:"Balance" xml:"Balance" require:"true"`
  Status *string `json:"Status" xml:"Status" require:"true"`
}

func (s QueryPrepaidCardsResponseDataPrepaidCard) String() string {
  return tea.Prettify(s)
}

func (s QueryPrepaidCardsResponseDataPrepaidCard) GoString() string {
  return s.String()
}

func (s *QueryPrepaidCardsResponseDataPrepaidCard) SetPrepaidCardId(v int64) *QueryPrepaidCardsResponseDataPrepaidCard {
  s.PrepaidCardId = &v
  return s
}

func (s *QueryPrepaidCardsResponseDataPrepaidCard) SetPrepaidCardNo(v string) *QueryPrepaidCardsResponseDataPrepaidCard {
  s.PrepaidCardNo = &v
  return s
}

func (s *QueryPrepaidCardsResponseDataPrepaidCard) SetGrantedTime(v string) *QueryPrepaidCardsResponseDataPrepaidCard {
  s.GrantedTime = &v
  return s
}

func (s *QueryPrepaidCardsResponseDataPrepaidCard) SetEffectiveTime(v string) *QueryPrepaidCardsResponseDataPrepaidCard {
  s.EffectiveTime = &v
  return s
}

func (s *QueryPrepaidCardsResponseDataPrepaidCard) SetExpiryTime(v string) *QueryPrepaidCardsResponseDataPrepaidCard {
  s.ExpiryTime = &v
  return s
}

func (s *QueryPrepaidCardsResponseDataPrepaidCard) SetApplicableProducts(v string) *QueryPrepaidCardsResponseDataPrepaidCard {
  s.ApplicableProducts = &v
  return s
}

func (s *QueryPrepaidCardsResponseDataPrepaidCard) SetApplicableScenarios(v string) *QueryPrepaidCardsResponseDataPrepaidCard {
  s.ApplicableScenarios = &v
  return s
}

func (s *QueryPrepaidCardsResponseDataPrepaidCard) SetNominalValue(v string) *QueryPrepaidCardsResponseDataPrepaidCard {
  s.NominalValue = &v
  return s
}

func (s *QueryPrepaidCardsResponseDataPrepaidCard) SetBalance(v string) *QueryPrepaidCardsResponseDataPrepaidCard {
  s.Balance = &v
  return s
}

func (s *QueryPrepaidCardsResponseDataPrepaidCard) SetStatus(v string) *QueryPrepaidCardsResponseDataPrepaidCard {
  s.Status = &v
  return s
}

type QueryCashCouponsRequest struct {
  ExpiryTimeEnd *string `json:"ExpiryTimeEnd" xml:"ExpiryTimeEnd"`
  ExpiryTimeStart *string `json:"ExpiryTimeStart" xml:"ExpiryTimeStart"`
  EffectiveOrNot *bool `json:"EffectiveOrNot" xml:"EffectiveOrNot"`
}

func (s QueryCashCouponsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCashCouponsRequest) GoString() string {
  return s.String()
}

func (s *QueryCashCouponsRequest) SetExpiryTimeEnd(v string) *QueryCashCouponsRequest {
  s.ExpiryTimeEnd = &v
  return s
}

func (s *QueryCashCouponsRequest) SetExpiryTimeStart(v string) *QueryCashCouponsRequest {
  s.ExpiryTimeStart = &v
  return s
}

func (s *QueryCashCouponsRequest) SetEffectiveOrNot(v bool) *QueryCashCouponsRequest {
  s.EffectiveOrNot = &v
  return s
}

type QueryCashCouponsResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryCashCouponsResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryCashCouponsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCashCouponsResponse) GoString() string {
  return s.String()
}

func (s *QueryCashCouponsResponse) SetRequestId(v string) *QueryCashCouponsResponse {
  s.RequestId = &v
  return s
}

func (s *QueryCashCouponsResponse) SetSuccess(v bool) *QueryCashCouponsResponse {
  s.Success = &v
  return s
}

func (s *QueryCashCouponsResponse) SetCode(v string) *QueryCashCouponsResponse {
  s.Code = &v
  return s
}

func (s *QueryCashCouponsResponse) SetMessage(v string) *QueryCashCouponsResponse {
  s.Message = &v
  return s
}

func (s *QueryCashCouponsResponse) SetData(v *QueryCashCouponsResponseData) *QueryCashCouponsResponse {
  s.Data = v
  return s
}

type QueryCashCouponsResponseData struct {
  CashCoupon []*QueryCashCouponsResponseDataCashCoupon `json:"CashCoupon" xml:"CashCoupon" require:"true" type:"Repeated"`
}

func (s QueryCashCouponsResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCashCouponsResponseData) GoString() string {
  return s.String()
}

func (s *QueryCashCouponsResponseData) SetCashCoupon(v []*QueryCashCouponsResponseDataCashCoupon) *QueryCashCouponsResponseData {
  s.CashCoupon = v
  return s
}

type QueryCashCouponsResponseDataCashCoupon struct     {
  CashCouponId *int64 `json:"CashCouponId" xml:"CashCouponId" require:"true"`
  CashCouponNo *string `json:"CashCouponNo" xml:"CashCouponNo" require:"true"`
  GrantedTime *string `json:"GrantedTime" xml:"GrantedTime" require:"true"`
  EffectiveTime *string `json:"EffectiveTime" xml:"EffectiveTime" require:"true"`
  ExpiryTime *string `json:"ExpiryTime" xml:"ExpiryTime" require:"true"`
  ApplicableProducts *string `json:"ApplicableProducts" xml:"ApplicableProducts" require:"true"`
  ApplicableScenarios *string `json:"ApplicableScenarios" xml:"ApplicableScenarios" require:"true"`
  NominalValue *string `json:"NominalValue" xml:"NominalValue" require:"true"`
  Balance *string `json:"Balance" xml:"Balance" require:"true"`
  Status *string `json:"Status" xml:"Status" require:"true"`
}

func (s QueryCashCouponsResponseDataCashCoupon) String() string {
  return tea.Prettify(s)
}

func (s QueryCashCouponsResponseDataCashCoupon) GoString() string {
  return s.String()
}

func (s *QueryCashCouponsResponseDataCashCoupon) SetCashCouponId(v int64) *QueryCashCouponsResponseDataCashCoupon {
  s.CashCouponId = &v
  return s
}

func (s *QueryCashCouponsResponseDataCashCoupon) SetCashCouponNo(v string) *QueryCashCouponsResponseDataCashCoupon {
  s.CashCouponNo = &v
  return s
}

func (s *QueryCashCouponsResponseDataCashCoupon) SetGrantedTime(v string) *QueryCashCouponsResponseDataCashCoupon {
  s.GrantedTime = &v
  return s
}

func (s *QueryCashCouponsResponseDataCashCoupon) SetEffectiveTime(v string) *QueryCashCouponsResponseDataCashCoupon {
  s.EffectiveTime = &v
  return s
}

func (s *QueryCashCouponsResponseDataCashCoupon) SetExpiryTime(v string) *QueryCashCouponsResponseDataCashCoupon {
  s.ExpiryTime = &v
  return s
}

func (s *QueryCashCouponsResponseDataCashCoupon) SetApplicableProducts(v string) *QueryCashCouponsResponseDataCashCoupon {
  s.ApplicableProducts = &v
  return s
}

func (s *QueryCashCouponsResponseDataCashCoupon) SetApplicableScenarios(v string) *QueryCashCouponsResponseDataCashCoupon {
  s.ApplicableScenarios = &v
  return s
}

func (s *QueryCashCouponsResponseDataCashCoupon) SetNominalValue(v string) *QueryCashCouponsResponseDataCashCoupon {
  s.NominalValue = &v
  return s
}

func (s *QueryCashCouponsResponseDataCashCoupon) SetBalance(v string) *QueryCashCouponsResponseDataCashCoupon {
  s.Balance = &v
  return s
}

func (s *QueryCashCouponsResponseDataCashCoupon) SetStatus(v string) *QueryCashCouponsResponseDataCashCoupon {
  s.Status = &v
  return s
}

type QueryAccountBalanceRequest struct {
}

func (s QueryAccountBalanceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountBalanceRequest) GoString() string {
  return s.String()
}

type QueryAccountBalanceResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *QueryAccountBalanceResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s QueryAccountBalanceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountBalanceResponse) GoString() string {
  return s.String()
}

func (s *QueryAccountBalanceResponse) SetRequestId(v string) *QueryAccountBalanceResponse {
  s.RequestId = &v
  return s
}

func (s *QueryAccountBalanceResponse) SetSuccess(v bool) *QueryAccountBalanceResponse {
  s.Success = &v
  return s
}

func (s *QueryAccountBalanceResponse) SetCode(v string) *QueryAccountBalanceResponse {
  s.Code = &v
  return s
}

func (s *QueryAccountBalanceResponse) SetMessage(v string) *QueryAccountBalanceResponse {
  s.Message = &v
  return s
}

func (s *QueryAccountBalanceResponse) SetData(v *QueryAccountBalanceResponseData) *QueryAccountBalanceResponse {
  s.Data = v
  return s
}

type QueryAccountBalanceResponseData struct {
  AvailableAmount *string `json:"AvailableAmount" xml:"AvailableAmount" require:"true"`
  AvailableCashAmount *string `json:"AvailableCashAmount" xml:"AvailableCashAmount" require:"true"`
  CreditAmount *string `json:"CreditAmount" xml:"CreditAmount" require:"true"`
  MybankCreditAmount *string `json:"MybankCreditAmount" xml:"MybankCreditAmount" require:"true"`
  Currency *string `json:"Currency" xml:"Currency" require:"true"`
}

func (s QueryAccountBalanceResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountBalanceResponseData) GoString() string {
  return s.String()
}

func (s *QueryAccountBalanceResponseData) SetAvailableAmount(v string) *QueryAccountBalanceResponseData {
  s.AvailableAmount = &v
  return s
}

func (s *QueryAccountBalanceResponseData) SetAvailableCashAmount(v string) *QueryAccountBalanceResponseData {
  s.AvailableCashAmount = &v
  return s
}

func (s *QueryAccountBalanceResponseData) SetCreditAmount(v string) *QueryAccountBalanceResponseData {
  s.CreditAmount = &v
  return s
}

func (s *QueryAccountBalanceResponseData) SetMybankCreditAmount(v string) *QueryAccountBalanceResponseData {
  s.MybankCreditAmount = &v
  return s
}

func (s *QueryAccountBalanceResponseData) SetCurrency(v string) *QueryAccountBalanceResponseData {
  s.Currency = &v
  return s
}

type DescribeResourcePackageProductRequest struct {
  ProductCode *string `json:"ProductCode" xml:"ProductCode"`
}

func (s DescribeResourcePackageProductRequest) String() string {
  return tea.Prettify(s)
}

func (s DescribeResourcePackageProductRequest) GoString() string {
  return s.String()
}

func (s *DescribeResourcePackageProductRequest) SetProductCode(v string) *DescribeResourcePackageProductRequest {
  s.ProductCode = &v
  return s
}

type DescribeResourcePackageProductResponse struct {
  RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
  OrderId *int64 `json:"OrderId" xml:"OrderId" require:"true"`
  Success *bool `json:"Success" xml:"Success" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Message *string `json:"Message" xml:"Message" require:"true"`
  Data *DescribeResourcePackageProductResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s DescribeResourcePackageProductResponse) String() string {
  return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponse) GoString() string {
  return s.String()
}

func (s *DescribeResourcePackageProductResponse) SetRequestId(v string) *DescribeResourcePackageProductResponse {
  s.RequestId = &v
  return s
}

func (s *DescribeResourcePackageProductResponse) SetOrderId(v int64) *DescribeResourcePackageProductResponse {
  s.OrderId = &v
  return s
}

func (s *DescribeResourcePackageProductResponse) SetSuccess(v bool) *DescribeResourcePackageProductResponse {
  s.Success = &v
  return s
}

func (s *DescribeResourcePackageProductResponse) SetCode(v string) *DescribeResourcePackageProductResponse {
  s.Code = &v
  return s
}

func (s *DescribeResourcePackageProductResponse) SetMessage(v string) *DescribeResourcePackageProductResponse {
  s.Message = &v
  return s
}

func (s *DescribeResourcePackageProductResponse) SetData(v *DescribeResourcePackageProductResponseData) *DescribeResourcePackageProductResponse {
  s.Data = v
  return s
}

type DescribeResourcePackageProductResponseData struct {
  ResourcePackages *DescribeResourcePackageProductResponseDataResourcePackages `json:"ResourcePackages" xml:"ResourcePackages" require:"true" type:"Struct"`
}

func (s DescribeResourcePackageProductResponseData) String() string {
  return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseData) GoString() string {
  return s.String()
}

func (s *DescribeResourcePackageProductResponseData) SetResourcePackages(v *DescribeResourcePackageProductResponseDataResourcePackages) *DescribeResourcePackageProductResponseData {
  s.ResourcePackages = v
  return s
}

type DescribeResourcePackageProductResponseDataResourcePackages struct {
  ResourcePackage []*DescribeResourcePackageProductResponseDataResourcePackagesResourcePackage `json:"ResourcePackage" xml:"ResourcePackage" require:"true" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseDataResourcePackages) String() string {
  return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseDataResourcePackages) GoString() string {
  return s.String()
}

func (s *DescribeResourcePackageProductResponseDataResourcePackages) SetResourcePackage(v []*DescribeResourcePackageProductResponseDataResourcePackagesResourcePackage) *DescribeResourcePackageProductResponseDataResourcePackages {
  s.ResourcePackage = v
  return s
}

type DescribeResourcePackageProductResponseDataResourcePackagesResourcePackage struct     {
  ProductCode *string `json:"ProductCode" xml:"ProductCode" require:"true"`
  ProductType *string `json:"ProductType" xml:"ProductType" require:"true"`
  Name *string `json:"Name" xml:"Name" require:"true"`
  PackageTypes *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypes `json:"PackageTypes" xml:"PackageTypes" require:"true" type:"Struct"`
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackage) String() string {
  return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackage) GoString() string {
  return s.String()
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackage) SetProductCode(v string) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackage {
  s.ProductCode = &v
  return s
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackage) SetProductType(v string) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackage {
  s.ProductType = &v
  return s
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackage) SetName(v string) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackage {
  s.Name = &v
  return s
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackage) SetPackageTypes(v *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypes) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackage {
  s.PackageTypes = v
  return s
}

type DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypes struct {
  PackageType []*DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageType `json:"PackageType" xml:"PackageType" require:"true" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypes) String() string {
  return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypes) GoString() string {
  return s.String()
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypes) SetPackageType(v []*DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageType) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypes {
  s.PackageType = v
  return s
}

type DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageType struct     {
  Name *string `json:"Name" xml:"Name" require:"true"`
  Code *string `json:"Code" xml:"Code" require:"true"`
  Properties *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties `json:"Properties" xml:"Properties" require:"true" type:"Struct"`
  Specifications *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications `json:"Specifications" xml:"Specifications" require:"true" type:"Struct"`
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageType) String() string {
  return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageType) GoString() string {
  return s.String()
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageType) SetName(v string) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageType {
  s.Name = &v
  return s
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageType) SetCode(v string) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageType {
  s.Code = &v
  return s
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageType) SetProperties(v *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageType {
  s.Properties = v
  return s
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageType) SetSpecifications(v *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageType {
  s.Specifications = v
  return s
}

type DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties struct {
  Property []*DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty `json:"Property" xml:"Property" require:"true" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties) String() string {
  return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties) GoString() string {
  return s.String()
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties) SetProperty(v []*DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties {
  s.Property = v
  return s
}

type DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty struct     {
  Name *string `json:"Name" xml:"Name" require:"true"`
  Value *string `json:"Value" xml:"Value" require:"true"`
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) String() string {
  return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) GoString() string {
  return s.String()
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) SetName(v string) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty {
  s.Name = &v
  return s
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) SetValue(v string) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty {
  s.Value = &v
  return s
}

type DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications struct {
  Specification []*DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification `json:"Specification" xml:"Specification" require:"true" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications) String() string {
  return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications) GoString() string {
  return s.String()
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications) SetSpecification(v []*DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications {
  s.Specification = v
  return s
}

type DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification struct     {
  Name *string `json:"Name" xml:"Name" require:"true"`
  Value *string `json:"Value" xml:"Value" require:"true"`
  AvailableDurations *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations `json:"AvailableDurations" xml:"AvailableDurations" require:"true" type:"Struct"`
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) String() string {
  return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) GoString() string {
  return s.String()
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) SetName(v string) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification {
  s.Name = &v
  return s
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) SetValue(v string) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification {
  s.Value = &v
  return s
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) SetAvailableDurations(v *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification {
  s.AvailableDurations = v
  return s
}

type DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations struct {
  AvailableDuration []*DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration `json:"AvailableDuration" xml:"AvailableDuration" require:"true" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations) String() string {
  return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations) GoString() string {
  return s.String()
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations) SetAvailableDuration(v []*DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations {
  s.AvailableDuration = v
  return s
}

type DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration struct     {
  Name *string `json:"Name" xml:"Name" require:"true"`
  Value *int `json:"Value" xml:"Value" require:"true"`
  Unit *string `json:"Unit" xml:"Unit" require:"true"`
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) String() string {
  return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) GoString() string {
  return s.String()
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) SetName(v string) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration {
  s.Name = &v
  return s
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) SetValue(v int) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration {
  s.Value = &v
  return s
}

func (s *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) SetUnit(v string) *DescribeResourcePackageProductResponseDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration {
  s.Unit = &v
  return s
}

type Client struct {
  rpc.Client
}

func NewClient(config *rpc.Config)(*Client, error) {
  client := new(Client)
  err := client.Init(config)
  return client, err
}

func (client *Client)Init(config *rpc.Config)(_err error) {
  _err = client.Client.Init(config  )
  if _err != nil {
    return _err
  }
  client.EndpointRule = tea.String("regional")
  client.EndpointMap = map[string]string{
    "ap-northeast-1": "business.ap-southeast-1.aliyuncs.com",
    "ap-northeast-2-pop": "business.ap-southeast-1.aliyuncs.com",
    "ap-south-1": "business.ap-southeast-1.aliyuncs.com",
    "ap-southeast-1": "business.ap-southeast-1.aliyuncs.com",
    "ap-southeast-2": "business.ap-southeast-1.aliyuncs.com",
    "ap-southeast-3": "business.ap-southeast-1.aliyuncs.com",
    "ap-southeast-5": "business.ap-southeast-1.aliyuncs.com",
    "cn-beijing": "business.aliyuncs.com",
    "cn-beijing-finance-1": "business.aliyuncs.com",
    "cn-beijing-finance-pop": "business.aliyuncs.com",
    "cn-beijing-gov-1": "business.aliyuncs.com",
    "cn-beijing-nu16-b01": "business.aliyuncs.com",
    "cn-chengdu": "business.aliyuncs.com",
    "cn-edge-1": "business.aliyuncs.com",
    "cn-fujian": "business.aliyuncs.com",
    "cn-haidian-cm12-c01": "business.aliyuncs.com",
    "cn-hangzhou": "business.aliyuncs.com",
    "cn-hangzhou-bj-b01": "business.aliyuncs.com",
    "cn-hangzhou-finance": "business.aliyuncs.com",
    "cn-hangzhou-internal-prod-1": "business.aliyuncs.com",
    "cn-hangzhou-internal-test-1": "business.aliyuncs.com",
    "cn-hangzhou-internal-test-2": "business.aliyuncs.com",
    "cn-hangzhou-internal-test-3": "business.aliyuncs.com",
    "cn-hangzhou-test-306": "business.aliyuncs.com",
    "cn-hongkong": "business.aliyuncs.com",
    "cn-hongkong-finance-pop": "business.aliyuncs.com",
    "cn-huhehaote": "business.aliyuncs.com",
    "cn-north-2-gov-1": "business.aliyuncs.com",
    "cn-qingdao": "business.aliyuncs.com",
    "cn-qingdao-nebula": "business.aliyuncs.com",
    "cn-shanghai": "business.aliyuncs.com",
    "cn-shanghai-et15-b01": "business.aliyuncs.com",
    "cn-shanghai-et2-b01": "business.aliyuncs.com",
    "cn-shanghai-finance-1": "business.aliyuncs.com",
    "cn-shanghai-inner": "business.aliyuncs.com",
    "cn-shanghai-internal-test-1": "business.aliyuncs.com",
    "cn-shenzhen": "business.aliyuncs.com",
    "cn-shenzhen-finance-1": "business.aliyuncs.com",
    "cn-shenzhen-inner": "business.aliyuncs.com",
    "cn-shenzhen-st4-d01": "business.aliyuncs.com",
    "cn-shenzhen-su18-b01": "business.aliyuncs.com",
    "cn-wuhan": "business.aliyuncs.com",
    "cn-yushanfang": "business.aliyuncs.com",
    "cn-zhangbei-na61-b01": "business.aliyuncs.com",
    "cn-zhangjiakou": "business.aliyuncs.com",
    "cn-zhangjiakou-na62-a01": "business.aliyuncs.com",
    "cn-zhengzhou-nebula-1": "business.aliyuncs.com",
    "eu-central-1": "business.ap-southeast-1.aliyuncs.com",
    "eu-west-1": "business.ap-southeast-1.aliyuncs.com",
    "eu-west-1-oxs": "business.ap-southeast-1.aliyuncs.com",
    "me-east-1": "business.ap-southeast-1.aliyuncs.com",
    "rus-west-1-pop": "business.ap-southeast-1.aliyuncs.com",
    "us-east-1": "business.ap-southeast-1.aliyuncs.com",
    "us-west-1": "business.ap-southeast-1.aliyuncs.com",
  }
  _err = client.CheckConfig(config)
  if _err != nil {
    return _err
  }
  client.Endpoint, _err = client.GetEndpoint(client.ProductId, client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
  if _err != nil {
    return _err
  }

  return nil
}



func (client *Client) QueryRIUtilizationDetailEx (request *QueryRIUtilizationDetailRequest, runtime *util.RuntimeOptions) (_result *QueryRIUtilizationDetailResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryRIUtilizationDetailResponse{}
  _body, _err := client.DoRequest(tea.String("QueryRIUtilizationDetail"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryRIUtilizationDetail (request *QueryRIUtilizationDetailRequest) (_result *QueryRIUtilizationDetailResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryRIUtilizationDetailResponse{}
  _body, _err := client.QueryRIUtilizationDetailEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryBillToOSSSubscriptionEx (request *QueryBillToOSSSubscriptionRequest, runtime *util.RuntimeOptions) (_result *QueryBillToOSSSubscriptionResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryBillToOSSSubscriptionResponse{}
  _body, _err := client.DoRequest(tea.String("QueryBillToOSSSubscription"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), nil, tea.ToMap(request), runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryBillToOSSSubscription (request *QueryBillToOSSSubscriptionRequest) (_result *QueryBillToOSSSubscriptionResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryBillToOSSSubscriptionResponse{}
  _body, _err := client.QueryBillToOSSSubscriptionEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryAccountBillEx (request *QueryAccountBillRequest, runtime *util.RuntimeOptions) (_result *QueryAccountBillResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryAccountBillResponse{}
  _body, _err := client.DoRequest(tea.String("QueryAccountBill"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryAccountBill (request *QueryAccountBillRequest) (_result *QueryAccountBillResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryAccountBillResponse{}
  _body, _err := client.QueryAccountBillEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) CreateCostUnitEx (request *CreateCostUnitRequest, runtime *util.RuntimeOptions) (_result *CreateCostUnitResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &CreateCostUnitResponse{}
  _body, _err := client.DoRequest(tea.String("CreateCostUnit"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) CreateCostUnit (request *CreateCostUnitRequest) (_result *CreateCostUnitResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &CreateCostUnitResponse{}
  _body, _err := client.CreateCostUnitEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) ModifyCostUnitEx (request *ModifyCostUnitRequest, runtime *util.RuntimeOptions) (_result *ModifyCostUnitResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &ModifyCostUnitResponse{}
  _body, _err := client.DoRequest(tea.String("ModifyCostUnit"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) ModifyCostUnit (request *ModifyCostUnitRequest) (_result *ModifyCostUnitResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &ModifyCostUnitResponse{}
  _body, _err := client.ModifyCostUnitEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryCostUnitEx (request *QueryCostUnitRequest, runtime *util.RuntimeOptions) (_result *QueryCostUnitResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryCostUnitResponse{}
  _body, _err := client.DoRequest(tea.String("QueryCostUnit"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryCostUnit (request *QueryCostUnitRequest) (_result *QueryCostUnitResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryCostUnitResponse{}
  _body, _err := client.QueryCostUnitEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) DeleteCostUnitEx (request *DeleteCostUnitRequest, runtime *util.RuntimeOptions) (_result *DeleteCostUnitResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &DeleteCostUnitResponse{}
  _body, _err := client.DoRequest(tea.String("DeleteCostUnit"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) DeleteCostUnit (request *DeleteCostUnitRequest) (_result *DeleteCostUnitResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &DeleteCostUnitResponse{}
  _body, _err := client.DeleteCostUnitEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) AllocateCostUnitResourceEx (request *AllocateCostUnitResourceRequest, runtime *util.RuntimeOptions) (_result *AllocateCostUnitResourceResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &AllocateCostUnitResourceResponse{}
  _body, _err := client.DoRequest(tea.String("AllocateCostUnitResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) AllocateCostUnitResource (request *AllocateCostUnitResourceRequest) (_result *AllocateCostUnitResourceResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &AllocateCostUnitResourceResponse{}
  _body, _err := client.AllocateCostUnitResourceEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryCostUnitResourceEx (request *QueryCostUnitResourceRequest, runtime *util.RuntimeOptions) (_result *QueryCostUnitResourceResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryCostUnitResourceResponse{}
  _body, _err := client.DoRequest(tea.String("QueryCostUnitResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryCostUnitResource (request *QueryCostUnitResourceRequest) (_result *QueryCostUnitResourceResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryCostUnitResourceResponse{}
  _body, _err := client.QueryCostUnitResourceEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) RenewResourcePackageEx (request *RenewResourcePackageRequest, runtime *util.RuntimeOptions) (_result *RenewResourcePackageResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &RenewResourcePackageResponse{}
  _body, _err := client.DoRequest(tea.String("RenewResourcePackage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) RenewResourcePackage (request *RenewResourcePackageRequest) (_result *RenewResourcePackageResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &RenewResourcePackageResponse{}
  _body, _err := client.RenewResourcePackageEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) UpgradeResourcePackageEx (request *UpgradeResourcePackageRequest, runtime *util.RuntimeOptions) (_result *UpgradeResourcePackageResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &UpgradeResourcePackageResponse{}
  _body, _err := client.DoRequest(tea.String("UpgradeResourcePackage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) UpgradeResourcePackage (request *UpgradeResourcePackageRequest) (_result *UpgradeResourcePackageResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &UpgradeResourcePackageResponse{}
  _body, _err := client.UpgradeResourcePackageEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) CreateAgAccountEx (request *CreateAgAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAgAccountResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &CreateAgAccountResponse{}
  _body, _err := client.DoRequest(tea.String("CreateAgAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) CreateAgAccount (request *CreateAgAccountRequest) (_result *CreateAgAccountResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &CreateAgAccountResponse{}
  _body, _err := client.CreateAgAccountEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) GetCustomerAccountInfoEx (request *GetCustomerAccountInfoRequest, runtime *util.RuntimeOptions) (_result *GetCustomerAccountInfoResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &GetCustomerAccountInfoResponse{}
  _body, _err := client.DoRequest(tea.String("GetCustomerAccountInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) GetCustomerAccountInfo (request *GetCustomerAccountInfoRequest) (_result *GetCustomerAccountInfoResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &GetCustomerAccountInfoResponse{}
  _body, _err := client.GetCustomerAccountInfoEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) GetCustomerListEx (request *GetCustomerListRequest, runtime *util.RuntimeOptions) (_result *GetCustomerListResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &GetCustomerListResponse{}
  _body, _err := client.DoRequest(tea.String("GetCustomerList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), nil, tea.ToMap(request), runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) GetCustomerList (request *GetCustomerListRequest) (_result *GetCustomerListResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &GetCustomerListResponse{}
  _body, _err := client.GetCustomerListEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) ChangeResellerConsumeAmountEx (request *ChangeResellerConsumeAmountRequest, runtime *util.RuntimeOptions) (_result *ChangeResellerConsumeAmountResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &ChangeResellerConsumeAmountResponse{}
  _body, _err := client.DoRequest(tea.String("ChangeResellerConsumeAmount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) ChangeResellerConsumeAmount (request *ChangeResellerConsumeAmountRequest) (_result *ChangeResellerConsumeAmountResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &ChangeResellerConsumeAmountResponse{}
  _body, _err := client.ChangeResellerConsumeAmountEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) SetResellerUserStatusEx (request *SetResellerUserStatusRequest, runtime *util.RuntimeOptions) (_result *SetResellerUserStatusResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &SetResellerUserStatusResponse{}
  _body, _err := client.DoRequest(tea.String("SetResellerUserStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) SetResellerUserStatus (request *SetResellerUserStatusRequest) (_result *SetResellerUserStatusResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &SetResellerUserStatusResponse{}
  _body, _err := client.SetResellerUserStatusEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) CreateResellerUserQuotaEx (request *CreateResellerUserQuotaRequest, runtime *util.RuntimeOptions) (_result *CreateResellerUserQuotaResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &CreateResellerUserQuotaResponse{}
  _body, _err := client.DoRequest(tea.String("CreateResellerUserQuota"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) CreateResellerUserQuota (request *CreateResellerUserQuotaRequest) (_result *CreateResellerUserQuotaResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &CreateResellerUserQuotaResponse{}
  _body, _err := client.CreateResellerUserQuotaEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) SetResellerUserQuotaEx (request *SetResellerUserQuotaRequest, runtime *util.RuntimeOptions) (_result *SetResellerUserQuotaResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &SetResellerUserQuotaResponse{}
  _body, _err := client.DoRequest(tea.String("SetResellerUserQuota"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) SetResellerUserQuota (request *SetResellerUserQuotaRequest) (_result *SetResellerUserQuotaResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &SetResellerUserQuotaResponse{}
  _body, _err := client.SetResellerUserQuotaEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryResellerAvailableQuotaEx (request *QueryResellerAvailableQuotaRequest, runtime *util.RuntimeOptions) (_result *QueryResellerAvailableQuotaResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryResellerAvailableQuotaResponse{}
  _body, _err := client.DoRequest(tea.String("QueryResellerAvailableQuota"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryResellerAvailableQuota (request *QueryResellerAvailableQuotaRequest) (_result *QueryResellerAvailableQuotaResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryResellerAvailableQuotaResponse{}
  _body, _err := client.QueryResellerAvailableQuotaEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) SetResellerUserAlarmThresholdEx (request *SetResellerUserAlarmThresholdRequest, runtime *util.RuntimeOptions) (_result *SetResellerUserAlarmThresholdResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &SetResellerUserAlarmThresholdResponse{}
  _body, _err := client.DoRequest(tea.String("SetResellerUserAlarmThreshold"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) SetResellerUserAlarmThreshold (request *SetResellerUserAlarmThresholdRequest) (_result *SetResellerUserAlarmThresholdResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &SetResellerUserAlarmThresholdResponse{}
  _body, _err := client.SetResellerUserAlarmThresholdEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryAccountTransactionsEx (request *QueryAccountTransactionsRequest, runtime *util.RuntimeOptions) (_result *QueryAccountTransactionsResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryAccountTransactionsResponse{}
  _body, _err := client.DoRequest(tea.String("QueryAccountTransactions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryAccountTransactions (request *QueryAccountTransactionsRequest) (_result *QueryAccountTransactionsResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryAccountTransactionsResponse{}
  _body, _err := client.QueryAccountTransactionsEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) UnsubscribeBillToOSSEx (request *UnsubscribeBillToOSSRequest, runtime *util.RuntimeOptions) (_result *UnsubscribeBillToOSSResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &UnsubscribeBillToOSSResponse{}
  _body, _err := client.DoRequest(tea.String("UnsubscribeBillToOSS"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) UnsubscribeBillToOSS (request *UnsubscribeBillToOSSRequest) (_result *UnsubscribeBillToOSSResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &UnsubscribeBillToOSSResponse{}
  _body, _err := client.UnsubscribeBillToOSSEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) SubscribeBillToOSSEx (request *SubscribeBillToOSSRequest, runtime *util.RuntimeOptions) (_result *SubscribeBillToOSSResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &SubscribeBillToOSSResponse{}
  _body, _err := client.DoRequest(tea.String("SubscribeBillToOSS"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) SubscribeBillToOSS (request *SubscribeBillToOSSRequest) (_result *SubscribeBillToOSSResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &SubscribeBillToOSSResponse{}
  _body, _err := client.SubscribeBillToOSSEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryUserOmsDataEx (request *QueryUserOmsDataRequest, runtime *util.RuntimeOptions) (_result *QueryUserOmsDataResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryUserOmsDataResponse{}
  _body, _err := client.DoRequest(tea.String("QueryUserOmsData"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryUserOmsData (request *QueryUserOmsDataRequest) (_result *QueryUserOmsDataResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryUserOmsDataResponse{}
  _body, _err := client.QueryUserOmsDataEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) CancelOrderEx (request *CancelOrderRequest, runtime *util.RuntimeOptions) (_result *CancelOrderResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &CancelOrderResponse{}
  _body, _err := client.DoRequest(tea.String("CancelOrder"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) CancelOrder (request *CancelOrderRequest) (_result *CancelOrderResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &CancelOrderResponse{}
  _body, _err := client.CancelOrderEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) ApplyInvoiceEx (request *ApplyInvoiceRequest, runtime *util.RuntimeOptions) (_result *ApplyInvoiceResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &ApplyInvoiceResponse{}
  _body, _err := client.DoRequest(tea.String("ApplyInvoice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) ApplyInvoice (request *ApplyInvoiceRequest) (_result *ApplyInvoiceResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &ApplyInvoiceResponse{}
  _body, _err := client.ApplyInvoiceEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryCustomerAddressListEx (request *QueryCustomerAddressListRequest, runtime *util.RuntimeOptions) (_result *QueryCustomerAddressListResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryCustomerAddressListResponse{}
  _body, _err := client.DoRequest(tea.String("QueryCustomerAddressList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryCustomerAddressList (request *QueryCustomerAddressListRequest) (_result *QueryCustomerAddressListResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryCustomerAddressListResponse{}
  _body, _err := client.QueryCustomerAddressListEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryEvaluateListEx (request *QueryEvaluateListRequest, runtime *util.RuntimeOptions) (_result *QueryEvaluateListResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryEvaluateListResponse{}
  _body, _err := client.DoRequest(tea.String("QueryEvaluateList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryEvaluateList (request *QueryEvaluateListRequest) (_result *QueryEvaluateListResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryEvaluateListResponse{}
  _body, _err := client.QueryEvaluateListEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryInvoicingCustomerListEx (request *QueryInvoicingCustomerListRequest, runtime *util.RuntimeOptions) (_result *QueryInvoicingCustomerListResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryInvoicingCustomerListResponse{}
  _body, _err := client.DoRequest(tea.String("QueryInvoicingCustomerList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryInvoicingCustomerList (request *QueryInvoicingCustomerListRequest) (_result *QueryInvoicingCustomerListResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryInvoicingCustomerListResponse{}
  _body, _err := client.QueryInvoicingCustomerListEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryBillOverviewEx (request *QueryBillOverviewRequest, runtime *util.RuntimeOptions) (_result *QueryBillOverviewResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryBillOverviewResponse{}
  _body, _err := client.DoRequest(tea.String("QueryBillOverview"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryBillOverview (request *QueryBillOverviewRequest) (_result *QueryBillOverviewResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryBillOverviewResponse{}
  _body, _err := client.QueryBillOverviewEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryBillEx (request *QueryBillRequest, runtime *util.RuntimeOptions) (_result *QueryBillResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryBillResponse{}
  _body, _err := client.DoRequest(tea.String("QueryBill"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryBill (request *QueryBillRequest) (_result *QueryBillResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryBillResponse{}
  _body, _err := client.QueryBillEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryInstanceBillEx (request *QueryInstanceBillRequest, runtime *util.RuntimeOptions) (_result *QueryInstanceBillResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryInstanceBillResponse{}
  _body, _err := client.DoRequest(tea.String("QueryInstanceBill"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryInstanceBill (request *QueryInstanceBillRequest) (_result *QueryInstanceBillResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryInstanceBillResponse{}
  _body, _err := client.QueryInstanceBillEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) EnableBillGenerationEx (request *EnableBillGenerationRequest, runtime *util.RuntimeOptions) (_result *EnableBillGenerationResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &EnableBillGenerationResponse{}
  _body, _err := client.DoRequest(tea.String("EnableBillGeneration"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) EnableBillGeneration (request *EnableBillGenerationRequest) (_result *EnableBillGenerationResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &EnableBillGenerationResponse{}
  _body, _err := client.EnableBillGenerationEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryRedeemEx (request *QueryRedeemRequest, runtime *util.RuntimeOptions) (_result *QueryRedeemResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryRedeemResponse{}
  _body, _err := client.DoRequest(tea.String("QueryRedeem"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryRedeem (request *QueryRedeemRequest) (_result *QueryRedeemResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryRedeemResponse{}
  _body, _err := client.QueryRedeemEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) ConvertChargeTypeEx (request *ConvertChargeTypeRequest, runtime *util.RuntimeOptions) (_result *ConvertChargeTypeResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &ConvertChargeTypeResponse{}
  _body, _err := client.DoRequest(tea.String("ConvertChargeType"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) ConvertChargeType (request *ConvertChargeTypeRequest) (_result *ConvertChargeTypeResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &ConvertChargeTypeResponse{}
  _body, _err := client.ConvertChargeTypeEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) CreateInstanceEx (request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &CreateInstanceResponse{}
  _body, _err := client.DoRequest(tea.String("CreateInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) CreateInstance (request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &CreateInstanceResponse{}
  _body, _err := client.CreateInstanceEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) ModifyInstanceEx (request *ModifyInstanceRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &ModifyInstanceResponse{}
  _body, _err := client.DoRequest(tea.String("ModifyInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) ModifyInstance (request *ModifyInstanceRequest) (_result *ModifyInstanceResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &ModifyInstanceResponse{}
  _body, _err := client.ModifyInstanceEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) DescribePricingModuleEx (request *DescribePricingModuleRequest, runtime *util.RuntimeOptions) (_result *DescribePricingModuleResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &DescribePricingModuleResponse{}
  _body, _err := client.DoRequest(tea.String("DescribePricingModule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) DescribePricingModule (request *DescribePricingModuleRequest) (_result *DescribePricingModuleResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &DescribePricingModuleResponse{}
  _body, _err := client.DescribePricingModuleEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryProductListEx (request *QueryProductListRequest, runtime *util.RuntimeOptions) (_result *QueryProductListResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryProductListResponse{}
  _body, _err := client.DoRequest(tea.String("QueryProductList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryProductList (request *QueryProductListRequest) (_result *QueryProductListResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryProductListResponse{}
  _body, _err := client.QueryProductListEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryInstanceGaapCostEx (request *QueryInstanceGaapCostRequest, runtime *util.RuntimeOptions) (_result *QueryInstanceGaapCostResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryInstanceGaapCostResponse{}
  _body, _err := client.DoRequest(tea.String("QueryInstanceGaapCost"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryInstanceGaapCost (request *QueryInstanceGaapCostRequest) (_result *QueryInstanceGaapCostResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryInstanceGaapCostResponse{}
  _body, _err := client.QueryInstanceGaapCostEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) RenewInstanceEx (request *RenewInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &RenewInstanceResponse{}
  _body, _err := client.DoRequest(tea.String("RenewInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) RenewInstance (request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &RenewInstanceResponse{}
  _body, _err := client.RenewInstanceEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) GetOrderDetailEx (request *GetOrderDetailRequest, runtime *util.RuntimeOptions) (_result *GetOrderDetailResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &GetOrderDetailResponse{}
  _body, _err := client.DoRequest(tea.String("GetOrderDetail"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) GetOrderDetail (request *GetOrderDetailRequest) (_result *GetOrderDetailResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &GetOrderDetailResponse{}
  _body, _err := client.GetOrderDetailEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryOrdersEx (request *QueryOrdersRequest, runtime *util.RuntimeOptions) (_result *QueryOrdersResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryOrdersResponse{}
  _body, _err := client.DoRequest(tea.String("QueryOrders"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryOrders (request *QueryOrdersRequest) (_result *QueryOrdersResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryOrdersResponse{}
  _body, _err := client.QueryOrdersEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryMonthlyInstanceConsumptionEx (request *QueryMonthlyInstanceConsumptionRequest, runtime *util.RuntimeOptions) (_result *QueryMonthlyInstanceConsumptionResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryMonthlyInstanceConsumptionResponse{}
  _body, _err := client.DoRequest(tea.String("QueryMonthlyInstanceConsumption"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryMonthlyInstanceConsumption (request *QueryMonthlyInstanceConsumptionRequest) (_result *QueryMonthlyInstanceConsumptionResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryMonthlyInstanceConsumptionResponse{}
  _body, _err := client.QueryMonthlyInstanceConsumptionEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QuerySettlementBillEx (request *QuerySettlementBillRequest, runtime *util.RuntimeOptions) (_result *QuerySettlementBillResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QuerySettlementBillResponse{}
  _body, _err := client.DoRequest(tea.String("QuerySettlementBill"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QuerySettlementBill (request *QuerySettlementBillRequest) (_result *QuerySettlementBillResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QuerySettlementBillResponse{}
  _body, _err := client.QuerySettlementBillEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryMonthlyBillEx (request *QueryMonthlyBillRequest, runtime *util.RuntimeOptions) (_result *QueryMonthlyBillResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryMonthlyBillResponse{}
  _body, _err := client.DoRequest(tea.String("QueryMonthlyBill"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryMonthlyBill (request *QueryMonthlyBillRequest) (_result *QueryMonthlyBillResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryMonthlyBillResponse{}
  _body, _err := client.QueryMonthlyBillEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) SetRenewalEx (request *SetRenewalRequest, runtime *util.RuntimeOptions) (_result *SetRenewalResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &SetRenewalResponse{}
  _body, _err := client.DoRequest(tea.String("SetRenewal"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) SetRenewal (request *SetRenewalRequest) (_result *SetRenewalResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &SetRenewalResponse{}
  _body, _err := client.SetRenewalEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryAvailableInstancesEx (request *QueryAvailableInstancesRequest, runtime *util.RuntimeOptions) (_result *QueryAvailableInstancesResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryAvailableInstancesResponse{}
  _body, _err := client.DoRequest(tea.String("QueryAvailableInstances"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryAvailableInstances (request *QueryAvailableInstancesRequest) (_result *QueryAvailableInstancesResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryAvailableInstancesResponse{}
  _body, _err := client.QueryAvailableInstancesEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) CreateResourcePackageEx (request *CreateResourcePackageRequest, runtime *util.RuntimeOptions) (_result *CreateResourcePackageResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &CreateResourcePackageResponse{}
  _body, _err := client.DoRequest(tea.String("CreateResourcePackage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) CreateResourcePackage (request *CreateResourcePackageRequest) (_result *CreateResourcePackageResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &CreateResourcePackageResponse{}
  _body, _err := client.CreateResourcePackageEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryResourcePackageInstancesEx (request *QueryResourcePackageInstancesRequest, runtime *util.RuntimeOptions) (_result *QueryResourcePackageInstancesResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryResourcePackageInstancesResponse{}
  _body, _err := client.DoRequest(tea.String("QueryResourcePackageInstances"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryResourcePackageInstances (request *QueryResourcePackageInstancesRequest) (_result *QueryResourcePackageInstancesResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryResourcePackageInstancesResponse{}
  _body, _err := client.QueryResourcePackageInstancesEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) GetResourcePackagePriceEx (request *GetResourcePackagePriceRequest, runtime *util.RuntimeOptions) (_result *GetResourcePackagePriceResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &GetResourcePackagePriceResponse{}
  _body, _err := client.DoRequest(tea.String("GetResourcePackagePrice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) GetResourcePackagePrice (request *GetResourcePackagePriceRequest) (_result *GetResourcePackagePriceResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &GetResourcePackagePriceResponse{}
  _body, _err := client.GetResourcePackagePriceEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) GetSubscriptionPriceEx (request *GetSubscriptionPriceRequest, runtime *util.RuntimeOptions) (_result *GetSubscriptionPriceResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &GetSubscriptionPriceResponse{}
  _body, _err := client.DoRequest(tea.String("GetSubscriptionPrice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) GetSubscriptionPrice (request *GetSubscriptionPriceRequest) (_result *GetSubscriptionPriceResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &GetSubscriptionPriceResponse{}
  _body, _err := client.GetSubscriptionPriceEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) GetPayAsYouGoPriceEx (request *GetPayAsYouGoPriceRequest, runtime *util.RuntimeOptions) (_result *GetPayAsYouGoPriceResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &GetPayAsYouGoPriceResponse{}
  _body, _err := client.DoRequest(tea.String("GetPayAsYouGoPrice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) GetPayAsYouGoPrice (request *GetPayAsYouGoPriceRequest) (_result *GetPayAsYouGoPriceResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &GetPayAsYouGoPriceResponse{}
  _body, _err := client.GetPayAsYouGoPriceEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryPrepaidCardsEx (request *QueryPrepaidCardsRequest, runtime *util.RuntimeOptions) (_result *QueryPrepaidCardsResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryPrepaidCardsResponse{}
  _body, _err := client.DoRequest(tea.String("QueryPrepaidCards"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryPrepaidCards (request *QueryPrepaidCardsRequest) (_result *QueryPrepaidCardsResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryPrepaidCardsResponse{}
  _body, _err := client.QueryPrepaidCardsEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryCashCouponsEx (request *QueryCashCouponsRequest, runtime *util.RuntimeOptions) (_result *QueryCashCouponsResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryCashCouponsResponse{}
  _body, _err := client.DoRequest(tea.String("QueryCashCoupons"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryCashCoupons (request *QueryCashCouponsRequest) (_result *QueryCashCouponsResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryCashCouponsResponse{}
  _body, _err := client.QueryCashCouponsEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) QueryAccountBalanceEx (request *QueryAccountBalanceRequest, runtime *util.RuntimeOptions) (_result *QueryAccountBalanceResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &QueryAccountBalanceResponse{}
  _body, _err := client.DoRequest(tea.String("QueryAccountBalance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), nil, tea.ToMap(request), runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) QueryAccountBalance (request *QueryAccountBalanceRequest) (_result *QueryAccountBalanceResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &QueryAccountBalanceResponse{}
  _body, _err := client.QueryAccountBalanceEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) DescribeResourcePackageProductEx (request *DescribeResourcePackageProductRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePackageProductResponse, _err error) {
  _err = util.ValidateModel(request)
  if _err != nil {
    return _result, _err
  }
  _result = &DescribeResourcePackageProductResponse{}
  _body, _err := client.DoRequest(tea.String("DescribeResourcePackageProduct"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-14"), tea.String("AK"), tea.ToMap(request), nil, runtime)
  if _err != nil {
    return _result, _err
  }
  _err = tea.Convert(_body, &_result)
  return _result, _err
}

func (client *Client) DescribeResourcePackageProduct (request *DescribeResourcePackageProductRequest) (_result *DescribeResourcePackageProductResponse, _err error) {
  runtime := &util.RuntimeOptions{}
  _result = &DescribeResourcePackageProductResponse{}
  _body, _err := client.DescribeResourcePackageProductEx(request, runtime)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func (client *Client) GetEndpoint (productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]string, endpoint *string) (_result *string, _err error) {
  if !tea.BoolValue(util.Empty(endpoint)) {
    _result = endpoint
    return _result , _err
  }

  if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(tea.String(endpointMap[tea.StringValue(regionId)]))) {
    _result = tea.String(endpointMap[tea.StringValue(regionId)])
    return _result, _err
  }

  _body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

