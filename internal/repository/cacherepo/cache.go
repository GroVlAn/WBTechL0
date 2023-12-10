package cacherepo

import (
	"encoding/json"
	"fmt"
	"github.com/GroVlAn/WBTechL0/internal/core"
	"github.com/GroVlAn/WBTechL0/internal/service"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	DCacheKey         = "delivery"
	OrdCacheKey       = "order"
	PmtCacheKey       = "payment"
	ProdCacheKey      = "product"
	defaultExpiration = 1 << 20 * time.Hour
	purgeTime         = 1 << 20 * time.Hour
	dateCreatedFormat = "2006-01-02T03:04:05Z"
)

type CacheRepo struct {
	log *logrus.Logger
	c   *cache.Cache
}

func NewCache(log *logrus.Logger) *CacheRepo {
	return &CacheRepo{
		log: log,
		c:   cache.New(defaultExpiration, purgeTime),
	}
}

func (ch *CacheRepo) SetDelivery(id int64, dRepr service.DeliveryRepr) {
	key := MakeKey(DCacheKey, strconv.FormatInt(id, 10))
	ch.c.Set(key, dRepr, cache.NoExpiration)
}

func (ch *CacheRepo) Delivery(id int64) (service.DeliveryRepr, error) {
	key := MakeKey(DCacheKey, strconv.FormatInt(id, 10))

	dCh, ok := ch.c.Get(key)

	if !ok {
		ch.log.Errorf("error not found delivery in cahce by id: %d", id)

		return service.DeliveryRepr{}, errors.New("delivery not found")
	}

	chBytes, errMarsh := json.Marshal(dCh)

	if errMarsh != nil {
		ch.log.Errorf("error can not marshal delivery cache data: %s", errMarsh.Error())

		return service.DeliveryRepr{}, errors.New("delivery not found")
	}

	var dRepr service.DeliveryRepr
	errUnMarsh := json.Unmarshal(chBytes, &dRepr)

	if errUnMarsh != nil {
		ch.log.Errorf("error can not unmarshal delivery cache data: %s", errUnMarsh.Error())

		return service.DeliveryRepr{}, errors.New("delivery not found")
	}

	return dRepr, nil
}

func (ch *CacheRepo) DeleteDelivery(id int64) {
	key := MakeKey(DCacheKey, strconv.FormatInt(id, 10))
	ch.c.Delete(key)
	ch.log.Infof("delete delivery from cache by id: %d", id)
}

func (ch *CacheRepo) SetOrder(ordUid string, ordRepr service.OrderRepr) {
	key := MakeKey(OrdCacheKey, ordUid)
	ch.c.Set(key, ordRepr, cache.NoExpiration)
}

func (ch *CacheRepo) Order(ordUid string) (service.OrderRepr, error) {
	key := MakeKey(OrdCacheKey, ordUid)

	ordCh, ok := ch.c.Get(key)

	if !ok {
		ch.log.Errorf("error not found order in cahce by order_uid: %s", ordUid)

		return service.OrderRepr{}, errors.New("order not found")
	}

	chBytes, errMarsh := json.Marshal(ordCh)

	if errMarsh != nil {
		ch.log.Errorf("error can not marshal order cache data: %s", errMarsh.Error())

		return service.OrderRepr{}, errors.New("order not found")
	}

	var ordRepr service.OrderRepr
	errUnMarsh := json.Unmarshal(chBytes, &ordRepr)

	if errUnMarsh != nil {
		ch.log.Errorf("error can not unmarshal order cache data: %s", errUnMarsh.Error())

		return service.OrderRepr{}, errors.New("order not found")
	}

	return ordRepr, nil
}

func (ch *CacheRepo) DeleteOrder(ordUid string) {
	key := MakeKey(OrdCacheKey, ordUid)
	ch.c.Delete(key)
	ch.log.Infof("delete order from cahce by order uid: %s", ordUid)
}

func (ch *CacheRepo) SetPayment(tran string, pmtRepr service.PaymentRepr) {
	key := MakeKey(PmtCacheKey, tran)
	ch.c.Set(key, pmtRepr, cache.NoExpiration)
}

func (ch *CacheRepo) Payment(tran string) (service.PaymentRepr, error) {
	key := MakeKey(PmtCacheKey, tran)

	pmtCh, ok := ch.c.Get(key)

	if !ok {
		ch.log.Errorf("error not found payment in cahce by transaction: %s", tran)

		return service.PaymentRepr{}, errors.New("payment not found")
	}

	chBytes, errMarsh := json.Marshal(pmtCh)

	if errMarsh != nil {
		ch.log.Errorf("error can not marshal payment cache data: %s", errMarsh.Error())

		return service.PaymentRepr{}, errors.New("payment not found")
	}

	var pmtRepr service.PaymentRepr
	errUnMarsh := json.Unmarshal(chBytes, &pmtRepr)

	if errUnMarsh != nil {
		ch.log.Errorf("error can not unmarshal payment cache data: %s", errUnMarsh.Error())

		return service.PaymentRepr{}, errors.New("payment not found")
	}

	return pmtRepr, nil
}

func (ch *CacheRepo) DeletePayment(tran string) {
	key := MakeKey(PmtCacheKey, tran)
	ch.c.Delete(key)
	ch.log.Infof("delete payment form cahce by transaction: %s", tran)
}

func (ch *CacheRepo) SetProduct(id int64, prodRepr service.ProductRepr) {
	key := MakeKey(ProdCacheKey, strconv.FormatInt(id, 10))
	ch.c.Set(key, prodRepr, defaultExpiration)
}

func (ch *CacheRepo) Product(id int64) (service.ProductRepr, error) {
	key := MakeKey(ProdCacheKey, strconv.FormatInt(id, 10))
	prodCh, ok := ch.c.Get(key)

	if !ok {
		ch.log.Errorf("error not found product in cahce by id: %s", id)

		return service.ProductRepr{}, errors.New("product not found")
	}

	chBytes, errMarsh := json.Marshal(prodCh)

	if errMarsh != nil {
		ch.log.Errorf("error can not marshal product cache data: %s", errMarsh.Error())

		return service.ProductRepr{}, errors.New("product not found")
	}

	var prodRepr service.ProductRepr
	errUnMarsh := json.Unmarshal(chBytes, &prodRepr)

	if errUnMarsh != nil {
		ch.log.Errorf("error can not unmarshal product cache data: %s", errUnMarsh.Error())

		return service.ProductRepr{}, errors.New("product not found")
	}

	return prodRepr, nil
}

func (ch *CacheRepo) DeleteProduct(id int64) {
	key := MakeKey(ProdCacheKey, strconv.FormatInt(id, 10))
	ch.c.Delete(key)
	ch.log.Infof("delete product from cahce by id: %d", id)
}

func (ch *CacheRepo) LoadAll(
	dAll func() ([]core.Delivery, error),
	ordAll func() ([]core.Order, error),
	pmtAll func() ([]core.Payment, error),
	prodAll func() ([]core.Product, error),
) {
	dlvs, errDlv := dAll()

	if errDlv == nil {
		for _, item := range dlvs {
			ch.SetDelivery(item.Id, service.DeliveryRepr(item))
		}
	}

	pmts, errPmts := pmtAll()

	if errPmts == nil {
		for _, item := range pmts {
			ch.SetPayment(item.Transaction, service.PaymentRepr(item))
		}
	}

	prods, errProds := prodAll()
	if errProds == nil {
		for _, item := range prods {
			ch.SetProduct(item.Id, service.ProductRepr(item))
		}
	}

	ords, errOrds := ordAll()

	if errOrds == nil {
		for _, item := range ords {

			dRepr, errDRepr := ch.Delivery(item.DeliveryId)

			if errDRepr != nil {
				continue
			}

			pmtRepr, errPmtRepr := ch.Payment(item.OrderUid)

			if errPmtRepr != nil {
				continue
			}

			prodsRepr := make([]service.ProductRepr, 0)

			for _, prod := range prods {
				if item.TrackNumber == prod.TrackNumber {
					prodsRepr = append(prodsRepr, service.ProductRepr(prod))
				}
			}

			if len(prodsRepr) == 0 {
				continue
			}

			ordRepr := service.OrderRepr{
				OrderUid:          item.OrderUid,
				TrackNumber:       item.TrackNumber,
				Entry:             item.Entry,
				Delivery:          dRepr,
				Payment:           pmtRepr,
				Items:             prodsRepr,
				Locale:            item.Locale,
				InternalSignature: item.InternalSignature,
				CustomerId:        item.CustomerId,
				DeliveryService:   item.DeliveryService,
				Shardkey:          item.Shardkey,
				SmId:              item.SmId,
				OofShard:          item.OofShard,
				DateCreated:       item.DateCreated.Format(dateCreatedFormat),
			}

			ch.SetOrder(item.OrderUid, ordRepr)
		}
	}
}

func MakeKey(addr string, id string) string {
	return fmt.Sprintf("%s-%s", addr, id)
}
