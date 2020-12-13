package UsersEntity
import (
	. "github.com/go-xe2/x/xf/ef/xq/xentity"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type UsersRec struct {
	UserId string	`json:"user_id,omitempty"`
	LoginName string	`json:"login_name,omitempty"`
	Pwd string	`json:"pwd,omitempty"`
	Enc string	`json:"enc,omitempty"`
	NickName string	`json:"nick_name,omitempty"`
	Name string	`json:"name,omitempty"`
	Sex int	`json:"sex,omitempty"`
	Mobile string	`json:"mobile,omitempty"`
	Qq string	`json:"qq,omitempty"`
	ProvinceId int	`json:"province_id,omitempty"`
	Province string	`json:"province,omitempty"`
	CityId int	`json:"city_id,omitempty"`
	City string	`json:"city,omitempty"`
	CountyId int	`json:"county_id,omitempty"`
	County string	`json:"county,omitempty"`
	TownId int	`json:"town_id,omitempty"`
	Town string	`json:"town,omitempty"`
	Address string	`json:"address,omitempty"`
	Head string	`json:"head,omitempty"`
	CrDate int64	`json:"cr_date,omitempty"`
	UpDate int64	`json:"up_date,omitempty"`
	Status int	`json:"status,omitempty"`
	FollowCount int	`json:"follow_count,omitempty"`
	VisitCount int	`json:"visit_count,omitempty"`
	PraiseCount int	`json:"praise_count,omitempty"`
	ComplaintCount int	`json:"complaint_count,omitempty"`
	AuditId int	`json:"audit_id,omitempty"`
	AuditSummery string	`json:"audit_summery,omitempty"`
	AuditDate int64	`json:"audit_date,omitempty"`
	IsExpert int16	`json:"is_expert,omitempty"`
	Lng float64	`json:"lng,omitempty"`
	Lat float64	`json:"lat,omitempty"`
	Region string	`json:"region,omitempty"`
	Nation string	`json:"nation,omitempty"`
	Birthday string	`json:"birthday,omitempty"`
	IdNumber string	`json:"id_number,omitempty"`
	Email string	`json:"email,omitempty"`
	Grade int	`json:"grade,omitempty"`
	ProductCount int	`json:"product_count,omitempty"`
	ProNdCount int	`json:"pro_nd_count,omitempty"`
	BaseCount int	`json:"base_count,omitempty"`
	BaseNdCount int	`json:"base_nd_count,omitempty"`
	CompanyCount int	`json:"company_count,omitempty"`
	FinancingCount int	`json:"financing_count,omitempty"`
	InvestmentCount int	`json:"investment_count,omitempty"`
}

type UsersEntity struct {
	*TEntity
	UserId EFString`ef:"@field(name='u_user_id',primary=true,alias='user_id');@dbType(type=varchar, size=30,allowNull=false)"`
	LoginName EFString`ef:"@field(name='u_login_name',alias='login_name');@dbType(type=varchar, size=30,allowNull=false)"`
	Pwd EFString`ef:"@field(name='u_pwd',alias='pwd');@dbType(type=varchar, size=50,allowNull=false)"`
	Enc EFString`ef:"@field(name='u_enc',alias='enc');@dbType(type=varchar, size=10,allowNull=false)"`
	NickName EFString`ef:"@field(name='u_nick_name',alias='nick_name');@dbType(type=varchar, size=128,allowNull=false)"`
	Name EFString`ef:"@field(name='u_name',alias='name');@dbType(type=varchar, size=60,allowNull=false)"`
	Sex EFInt`ef:"@field(name='u_sex',alias='sex');@dbType(type=int,allowNull=false)"`
	Mobile EFString`ef:"@field(name='u_mobile',alias='mobile');@dbType(type=varchar, size=30,allowNull=false)"`
	Qq EFString`ef:"@field(name='u_qq',alias='qq');@dbType(type=varchar, size=30,allowNull=false)"`
	ProvinceId EFInt`ef:"@field(name='u_province_id',alias='province_id');@dbType(type=int,allowNull=false,default=0)"`
	Province EFString`ef:"@field(name='u_province',alias='province');@dbType(type=varchar, size=80,allowNull=false)"`
	CityId EFInt`ef:"@field(name='u_city_id',alias='city_id');@dbType(type=int,allowNull=false,default=0)"`
	City EFString`ef:"@field(name='u_city',alias='city');@dbType(type=varchar, size=80,allowNull=false)"`
	CountyId EFInt`ef:"@field(name='u_county_id',alias='county_id');@dbType(type=int,allowNull=false,default=0)"`
	County EFString`ef:"@field(name='u_county',alias='county');@dbType(type=varchar, size=80,allowNull=false)"`
	TownId EFInt`ef:"@field(name='u_town_id',alias='town_id');@dbType(type=int,allowNull=false,default=0)"`
	Town EFString`ef:"@field(name='u_town',alias='town');@dbType(type=varchar, size=80,allowNull=false)"`
	Address EFString`ef:"@field(name='u_address',alias='address');@dbType(type=varchar, size=255,allowNull=false)"`
	Head EFString`ef:"@field(name='u_head',alias='head');@dbType(type=varchar, size=255,allowNull=false)"`
	CrDate EFInt64`ef:"@field(name='u_cr_date',alias='cr_date');@dbType(type=bigint,allowNull=false,default=0)"`
	UpDate EFInt64`ef:"@field(name='u_up_date',alias='up_date');@dbType(type=bigint,allowNull=false,default=0)"`
	Status EFInt`ef:"@field(name='u_status',alias='status');@dbType(type=int,allowNull=false,default=0)"`
	FollowCount EFInt`ef:"@field(name='u_follow_count',alias='follow_count');@dbType(type=int,allowNull=false,default=0)"`
	VisitCount EFInt`ef:"@field(name='u_visit_count',alias='visit_count');@dbType(type=int,allowNull=false,default=0)"`
	PraiseCount EFInt`ef:"@field(name='u_praise_count',alias='praise_count');@dbType(type=int,allowNull=false,default=0)"`
	ComplaintCount EFInt`ef:"@field(name='u_complaint_count',alias='complaint_count');@dbType(type=int,allowNull=false,default=0)"`
	AuditId EFInt`ef:"@field(name='u_audit_id',alias='audit_id');@dbType(type=int,allowNull=false,default=0)"`
	AuditSummery EFString`ef:"@field(name='u_audit_summery',alias='audit_summery');@dbType(type=varchar, size=255,allowNull=false)"`
	AuditDate EFInt64`ef:"@field(name='u_audit_date',alias='audit_date');@dbType(type=bigint,allowNull=false,default=0)"`
	IsExpert EFInt16`ef:"@field(name='u_is_expert',alias='is_expert');@dbType(type=tinyint,allowNull=false,default=0)"`
	Lng EFDouble`ef:"@field(name='u_lng',alias='lng');@dbType(type=decimal,size=10,decimal=6,allowNull=false,default=0.000000)"`
	Lat EFDouble`ef:"@field(name='u_lat',alias='lat');@dbType(type=decimal,size=10,decimal=6,allowNull=false,default=0.000000)"`
	Region EFString`ef:"@field(name='u_region',alias='region');@dbType(type=varchar, size=200,allowNull=false)"`
	Nation EFString`ef:"@field(name='u_nation',alias='nation');@dbType(type=varchar, size=20,allowNull=false)"`
	Birthday EFString`ef:"@field(name='u_birthday',alias='birthday');@dbType(type=varchar, size=10,allowNull=false)"`
	IdNumber EFString`ef:"@field(name='u_id_number',alias='id_number');@dbType(type=varchar, size=20,allowNull=false)"`
	Email EFString`ef:"@field(name='u_email',alias='email');@dbType(type=varchar, size=30,allowNull=false)"`
	Grade EFInt`ef:"@field(name='u_grade',alias='grade');@dbType(type=int,allowNull=false,default=0)"`
	ProductCount EFInt`ef:"@field(name='u_product_count',alias='product_count');@dbType(type=int,allowNull=false,default=0)"`
	ProNdCount EFInt`ef:"@field(name='u_pro_nd_count',alias='pro_nd_count');@dbType(type=int,allowNull=false,default=0)"`
	BaseCount EFInt`ef:"@field(name='u_base_count',alias='base_count');@dbType(type=int,allowNull=false,default=0)"`
	BaseNdCount EFInt`ef:"@field(name='u_base_nd_count',alias='base_nd_count');@dbType(type=int,allowNull=false,default=0)"`
	CompanyCount EFInt`ef:"@field(name='u_company_count',alias='company_count');@dbType(type=int,allowNull=false,default=0)"`
	FinancingCount EFInt`ef:"@field(name='u_financing_count',alias='financing_count');@dbType(type=int,allowNull=false,default=0)"`
	InvestmentCount EFInt`ef:"@field(name='u_investment_count',alias='investment_count');@dbType(type=int,allowNull=false,default=0)"`
}

var UsersEntityClass = ClassOfEntity(func() Entity {
	var inst Entity = new(UsersEntity)
	return inst
}, []XqAttribute{MakeEntityAttr("users", "uss")})

// 设置实体继承类实例, 供构架内部调用
func (ent *UsersEntity) Implement(supper interface{}) {
	if v, ok := supper.(*TEntity); ok {
		ent.TEntity = v
	}
}

// 继承的父类
func (ent *UsersEntity) Supper() Entity {
	return ent.TEntity
}

// 实体构造方法
func (ent *UsersEntity) Constructor(attrs []XqAttribute, inherited ...interface{}) interface{} {
	ent.Supper().Constructor(attrs, inherited...)
	return ent
}

func (ent *UsersEntity) String() string {
	return "Users"
}

func (ent *UsersEntity) Record() *UsersRec {
	if !ent.IsOpen() {
		return &UsersRec{}
	}
	result := &UsersRec{}
	if v, ok := ent.UserId.TryStr(); ok {
		result.UserId = v
	}
	if v, ok := ent.LoginName.TryStr(); ok {
		result.LoginName = v
	}
	if v, ok := ent.Pwd.TryStr(); ok {
		result.Pwd = v
	}
	if v, ok := ent.Enc.TryStr(); ok {
		result.Enc = v
	}
	if v, ok := ent.NickName.TryStr(); ok {
		result.NickName = v
	}
	if v, ok := ent.Name.TryStr(); ok {
		result.Name = v
	}
	if v, ok := ent.Sex.TryInt(); ok {
		result.Sex = v
	}
	if v, ok := ent.Mobile.TryStr(); ok {
		result.Mobile = v
	}
	if v, ok := ent.Qq.TryStr(); ok {
		result.Qq = v
	}
	if v, ok := ent.ProvinceId.TryInt(); ok {
		result.ProvinceId = v
	}
	if v, ok := ent.Province.TryStr(); ok {
		result.Province = v
	}
	if v, ok := ent.CityId.TryInt(); ok {
		result.CityId = v
	}
	if v, ok := ent.City.TryStr(); ok {
		result.City = v
	}
	if v, ok := ent.CountyId.TryInt(); ok {
		result.CountyId = v
	}
	if v, ok := ent.County.TryStr(); ok {
		result.County = v
	}
	if v, ok := ent.TownId.TryInt(); ok {
		result.TownId = v
	}
	if v, ok := ent.Town.TryStr(); ok {
		result.Town = v
	}
	if v, ok := ent.Address.TryStr(); ok {
		result.Address = v
	}
	if v, ok := ent.Head.TryStr(); ok {
		result.Head = v
	}
	if v, ok := ent.CrDate.TryInt64(); ok {
		result.CrDate = v
	}
	if v, ok := ent.UpDate.TryInt64(); ok {
		result.UpDate = v
	}
	if v, ok := ent.Status.TryInt(); ok {
		result.Status = v
	}
	if v, ok := ent.FollowCount.TryInt(); ok {
		result.FollowCount = v
	}
	if v, ok := ent.VisitCount.TryInt(); ok {
		result.VisitCount = v
	}
	if v, ok := ent.PraiseCount.TryInt(); ok {
		result.PraiseCount = v
	}
	if v, ok := ent.ComplaintCount.TryInt(); ok {
		result.ComplaintCount = v
	}
	if v, ok := ent.AuditId.TryInt(); ok {
		result.AuditId = v
	}
	if v, ok := ent.AuditSummery.TryStr(); ok {
		result.AuditSummery = v
	}
	if v, ok := ent.AuditDate.TryInt64(); ok {
		result.AuditDate = v
	}
	if v, ok := ent.IsExpert.TryInt16(); ok {
		result.IsExpert = v
	}
	if v, ok := ent.Lng.TryDouble(); ok {
		result.Lng = v
	}
	if v, ok := ent.Lat.TryDouble(); ok {
		result.Lat = v
	}
	if v, ok := ent.Region.TryStr(); ok {
		result.Region = v
	}
	if v, ok := ent.Nation.TryStr(); ok {
		result.Nation = v
	}
	if v, ok := ent.Birthday.TryStr(); ok {
		result.Birthday = v
	}
	if v, ok := ent.IdNumber.TryStr(); ok {
		result.IdNumber = v
	}
	if v, ok := ent.Email.TryStr(); ok {
		result.Email = v
	}
	if v, ok := ent.Grade.TryInt(); ok {
		result.Grade = v
	}
	if v, ok := ent.ProductCount.TryInt(); ok {
		result.ProductCount = v
	}
	if v, ok := ent.ProNdCount.TryInt(); ok {
		result.ProNdCount = v
	}
	if v, ok := ent.BaseCount.TryInt(); ok {
		result.BaseCount = v
	}
	if v, ok := ent.BaseNdCount.TryInt(); ok {
		result.BaseNdCount = v
	}
	if v, ok := ent.CompanyCount.TryInt(); ok {
		result.CompanyCount = v
	}
	if v, ok := ent.FinancingCount.TryInt(); ok {
		result.FinancingCount = v
	}
	if v, ok := ent.InvestmentCount.TryInt(); ok {
		result.InvestmentCount = v
	}
	return result
}
