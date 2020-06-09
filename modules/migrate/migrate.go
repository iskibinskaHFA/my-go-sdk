package migrate

import "github.com/jinzhu/gorm"

//Migrate performs basic migration for test environment
func Migrate(db *gorm.DB) {
	db.DropTable(&CalcStepsLog{})
	db.DropTable(&CalcStepsLogDefinition{})
	db.DropTable(&RateDefinition{})
	db.DropTable(&RateFormula{})
	db.DropTable(&Rate{})
	db.DropTable(&Release{})
	db.DropTable(&WorksWriter{})
	db.DropTable(&Writer{})
	db.DropTable(&Resource{})
	db.DropTable(&UsageSummary{})
	db.DropTable(&Header{})
	db.DropTable(&Service{})
	db.DropTable(&Work{})

	db.CreateTable(&Header{})
	db.CreateTable(&Release{})
	db.CreateTable(&Resource{})
	db.CreateTable(&Work{})
	db.CreateTable(&Writer{})
	db.CreateTable(&WorksWriter{})

	db.Model(Resource{}).AddForeignKey("origin_id", "usage.headers(header_id)", "CASCADE", "CASCADE")
	db.Model(Resource{}).AddForeignKey("work_id", "usage.works(work_id)", "CASCADE", "CASCADE")
	db.Model(WorksWriter{}).AddForeignKey("work_id", "usage.works(work_id)", "CASCADE", "CASCADE")
	db.Model(WorksWriter{}).AddForeignKey("party_id", "usage.writers(party_id)", "CASCADE", "CASCADE")

	db.CreateTable(&RateFormula{})
	db.CreateTable(&Service{})
	db.CreateTable(&UsageSummary{})
	db.CreateTable(&RateDefinition{})
	db.CreateTable(&Rate{})
	db.CreateTable(&CalcStepsLogDefinition{})
	db.CreateTable(&CalcStepsLog{})

	db.Model(UsageSummary{}).AddForeignKey("service_id", "royalty.services(service_id)", "CASCADE", "CASCADE")
	db.Model(UsageSummary{}).AddForeignKey("header_id", "usage.headers(header_id)", "CASCADE", "CASCADE")
	db.Model(RateDefinition{}).AddForeignKey("service_id", "royalty.services(service_id)", "CASCADE", "CASCADE")
	db.Model(RateDefinition{}).AddForeignKey("rate_formula_id", "royalty.rate_formulas(rate_formula_id)", "CASCADE", "CASCADE")
	db.Model(Rate{}).AddForeignKey("usage_summary_id", "usage.usage_summaries(usage_summary_id)", "CASCADE", "CASCADE")
	db.Model(CalcStepsLogDefinition{}).AddForeignKey("service_id", "royalty.services(service_id)", "CASCADE", "CASCADE")
	db.Model(CalcStepsLog{}).AddForeignKey("usage_summary_id", "usage.usage_summaries(usage_summary_id)", "CASCADE", "CASCADE")
	db.Model(CalcStepsLog{}).AddForeignKey("log_definition_id", "royalty.calc_steps_log_definitions(log_definition_id)", "CASCADE",
		"CASCADE")
}
