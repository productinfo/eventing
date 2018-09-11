package main

import (
	"fmt"
	"math/rand"
)

type Merchant struct {
	Type string `json:"type"`
	ID   string `json:"merchantid"`
	Name string `json:"name"`
	City City   `json:"city"`
}

var m_counter = 0

func MakeMerchant() Merchant {
	name := merchants[rand.Intn(len(merchants))]
	suffix := suffixes[rand.Intn(len(suffixes))]
	city := Cities[rand.Intn(len(Cities))]
	id := fmt.Sprintf("merchant-%v-%v", rand.Int(), m_counter)
	m_counter++
	merchant := Merchant{"merchant", id, name + " " + suffix, city}
	return merchant
}

var suffixes = []string{
	"Inc",
	"Corp",
	"Company",
	"Corporation",
	"Incorporated",
	"LLC",
}

var merchants = []string{
	`3 Round Stones,`,
	`48 Factoring`,
	`5PSolutions`,
	`Abt Associates`,
	`Accela`,
	`Accenture`,
	`AccuWeather`,
	`Acxiom`,
	`Adaptive`,
	`Adobe Digital Government`,
	`Aidin`,
	`Alarm.com`,
	`Allianz`,
	`Allied Van Lines`,
	`AllState Insurance Group`,
	`Alltuition`,
	`Altova`,
	`Amazon Web Services`,
	`American Red Ball Movers`,
	`Amida Technology Solutions`,
	`Analytica`,
	`Apextech`,
	`Appallicious`,
	`Aquicore`,
	`Archimedes`,
	`AreaVibes`,
	`Arpin Van Lines`,
	`Arrive Labs`,
	`ASC Partners`,
	`Asset4`,
	`Atlas Van Lines`,
	`AtSite`,
	`Aunt Bertha,`,
	`Aureus Sciences`,
	`AutoGrid Systems`,
	`Avalara`,
	`Avvo`,
	`Ayasdi`,
	`Azavea`,
	`BaleFire Global`,
	`Barchart`,
	`Be Informed`,
	`Bekins`,
	`Berkery Noyes MandASoft`,
	`Berkshire Hathaway`,
	`BetterLesson`,
	`BillGuard`,
	`Bing`,
	`Biovia`,
	`BizVizz`,
	`BlackRock`,
	`Bloomberg`,
	`Booz Allen Hamilton`,
	`Boston Consulting Group`,
	`Boundless`,
	`Bridgewater`,
	`Brightscope`,
	`BuildFax`,
	`Buildingeye`,
	`BuildZoom`,
	`Business and Legal Resources`,
	`Business Monitor International`,
	`Calcbench,`,
	`Cambridge Information Group`,
	`Cambridge Semantics`,
	`CAN Capital`,
	`Canon`,
	`Capital Cube`,
	`Cappex`,
	`Captricity`,
	`CareSet Systems`,
	`Careset.com`,
	`CARFAX`,
	`Caspio`,
	`Castle Biosciences`,
	`CB Insights`,
	`Ceiba Solutions`,
	`Center for Responsive Politics`,
	`Cerner`,
	`Certara`,
	`CGI`,
	`Charles River Associates`,
	`Charles Schwab`,
	`Chemical Abstracts Service`,
	`Child Care Desk`,
	`Chubb`,
	`Citigroup`,
	`CityScan`,
	`CitySourced`,
	`Civic Impulse`,
	`Civic Insight`,
	`Civinomics`,
	`Civis Analytics`,
	`Clean Power Finance`,
	`ClearHealthCosts`,
	`ClearStory Data`,
	`Climate`,
	`CliniCast`,
	`Cloudmade`,
	`Cloudspyre`,
	`Code for America`,
	`Code-N`,
	`Collective IP`,
	`College Abacus, an ECMC initiative`,
	`College Board`,
	`Compared Care`,
	`Compendia Bioscience Life Technologies`,
	`Compliance and Risks`,
	`Computer Packages`,
	`CONNECT-DOT `,
	`ConnectEDU`,
	`Connotate`,
	`Construction Monitor`,
	`Consumer Reports`,
	`CoolClimate`,
	`Copyright Clearance Center`,
	`CoreLogic`,
	`CostQuest`,
	`Credit Karma`,
	`Credit Sesame`,
	`CrowdANALYTIX`,
	`Dabo Health`,
	`DataLogix`,
	`DataMade`,
	`DataMarket`,
	`Datamyne`,
	`DataWeave`,
	`Deloitte`,
	`DemystData`,
	`Department of Better Technology`,
	`Development Seed`,
	`Docket Alarm,`,
	`Dow Jones & Co`,
	`Dun & Bradstreet`,
	`Earth Networks`,
	`EarthObserver App`,
	`Earthquake Alert`,
	`Eat Shop Sleep`,
	`Ecodesk`,
	`eInstitutional`,
	`Embark`,
	`EMC`,
	`Energy Points,`,
	`Energy Solutions Forum`,
	`Enervee`,
	`Enigma.io`,
	`Ensco`,
	`Environmental Data Resources`,
	`Epsilon`,
	`Equal Speed`,
	`Equifax`,
	`Equilar`,
	`Ernst & Young LLP`,
	`eScholar `,
	`Esri`,
	`Estately`,
	`Everyday Health`,
	`Evidera`,
	`Experian`,
	`Expert Health Data Programming,`,
	`Exversion`,
	`Ez-XBRL`,
	`Factset`,
	`Factual`,
	`Farmers`,
	`FarmLogs`,
	`Fastcase`,
	`Fidelity Investments`,
	`FindTheBest.com`,
	`First Fuel Software`,
	`FirstPoint,`,
	`Fitch`,
	`FlightAware`,
	`FlightStats`,
	`FlightView`,
	`Food+Tech Connect`,
	`Forrester Research`,
	`Foursquare`,
	`Fujitsu`,
	`Funding Circle`,
	`FutureAdvisor`,
	`Fuzion Apps,`,
	`Gallup`,
	`Galorath orporated`,
	`Garmin`,
	`Genability`,
	`GenoSpace`,
	`Geofeedia`,
	`Geolytics`,
	`Geoscape`,
	`GetRaised`,
	`GitHub`,
	`Glassy Media`,
	`Golden Helix`,
	`GoodGuide`,
	`Google Maps`,
	`Google Public Data Explorer`,
	`Government Transaction Services`,
	`Govini`,
	`GovTribe`,
	`Govzilla,`,
	`gRadiant Research`,
	`Graebel Van Lines`,
	`Graematter,`,
	`Granicus`,
	`GreatSchools`,
	`GuideStar`,
	`H3 Biomedicine`,
	`Harris`,
	`HDScores,`,
	`Headlight`,
	`Healthgrades`,
	`Healthline`,
	`HealthMap`,
	`HealthPocket,`,
	`HelloWallet`,
	`HERE`,
	`Honest Buildings`,
	`HopStop`,
	`Housefax`,
	`How's My Offer?`,
	`IBM`,
	`ideas42`,
	`iFactor Consulting`,
	`IFI CLAIMS Patent Services`,
	`iMedicare`,
	`Impact Forecasting`,
	`Impaq International`,
	`Import.io`,
	`IMS Health`,
	`InCadence`,
	`indoo.rs`,
	`InfoCommerce Group`,
	`Informatica`,
	`InnoCentive`,
	`Innography`,
	`Innovest Systems`,
	`Inovalon`,
	`Inrix Traffic`,
	`Intelius`,
	`Intermap Technologies`,
	`Investormill`,
	`Iodine`,
	`IPHIX`,
	`iRecycle`,
	`iTriage`,
	`IVES Group`,
	`IW Financial`,
	`JJ Keller`,
	`J.P. Morgan Chase`,
	`Junar,`,
	`Junyo`,
	`Jurispect`,
	`Kaiser Permanante`,
	`karmadata`,
	`Keychain Logistics`,
	`KidAdmit,`,
	`Kimono Labs`,
	`KLD Research`,
	`Knoema`,
	`Knowledge Agency`,
	`KPMG`,
	`Kroll Bond Ratings Agency`,
	`Kyruus`,
	`Lawdragon`,
	`Legal Science Partners`,
	`Cyte`,
	`LegiNation,`,
	`LegiStorm`,
	`Lenddo`,
	`Lending Club`,
	`Level One Technologies`,
	`LexisNexis`,
	`Liberty Mutual Insurance Cos`,
	`Lilly Open Innovation Drug Discovery`,
	`Liquid Robotics`,
	`Locavore`,
	`LOGIXDATA,`,
	`LoopNet`,
	`Loqate,`,
	`LoseIt.com`,
	`LOVELAND Technologies`,
	`Lucid`,
	`Lumesis,`,
	`Mango Transit`,
	`Mapbox`,
	`Maponics`,
	`MapQuest`,
	`Marinexplore,`,
	`MarketSense`,
	`Marlin & Associates`,
	`Marlin Alter and Associates`,
	`McGraw Hill Financial`,
	`McKinsey`,
	`MedWatcher`,
	`Mercaris`,
	`Merrill`,
	`Merrill Lynch`,
	`MetLife`,
	`mHealthCoach`,
	`MicroBilt`,
	`Microsoft Windows Azure Marketplace`,
	`Mint`,
	`Moody's`,
	`Morgan Stanley`,
	`Morningstar,`,
	`Mozio`,
	`MuckRock.com`,
	`Munetrix`,
	`Municode`,
	`National Van Lines`,
	`Nationwide Mutual Insurance Company`,
	`Nautilytics`,
	`Navico`,
	`NERA Economic Consulting`,
	`NerdWallet`,
	`New Media Parents`,
	`Next Step Living`,
	`NextBus`,
	`nGAP orporated`,
	`Nielsen`,
	`Noesis`,
	`NonprofitMetrics`,
	`North American Van Lines`,
	`Noveda Technologies`,
	`NuCivic`,
	`Numedii`,
	`Oliver Wyman`,
	`OnDeck`,
	`OnStar`,
	`Ontodia,`,
	`Onvia`,
	`Open Data Nation`,
	`OpenCounter`,
	`OpenGov`,
	`OpenPlans`,
	`OpportunitySpace,`,
	`Optensity`,
	`optiGov`,
	`OptumInsight`,
	`Orlin Research`,
	`OSIsoft`,
	`OTC Markets`,
	`Outline`,
	`Oversight Systems`,
	`Overture Technologies`,
	`Owler`,
	`Palantir Technologies`,
	`Panjiva`,
	`Parsons Brinckerhoff`,
	`Patently-O`,
	`PatientsLikeMe`,
	`Pave`,
	`Paxata`,
	`PayScale,`,
	`PeerJ`,
	`People Power`,
	`Persint`,
	`Personal Democracy Media`,
	`Personal,`,
	`Personalis`,
	`Peterson's`,
	`PEV4me.com`,
	`PIXIA`,
	`PlaceILive.com`,
	`PlanetEcosystems`,
	`PlotWatt`,
	`Plus-U`,
	`PolicyMap`,
	`Politify`,
	`Poncho App`,
	`POPVOX`,
	`Porch`,
	`PossibilityU`,
	`PowerAdvocate`,
	`Practice Fusion`,
	`Predilytics`,
	`PricewaterhouseCoopers`,
	`ProgrammableWeb`,
	`Progressive Insurance Group`,
	`Propeller Health`,
	`ProPublica`,
	`PublicEngines`,
	`PYA Analytics`,
	`Qado Energy,`,
	`Quandl`,
	`Quertle`,
	`Quid`,
	`R R Donnelley`,
	`RAND`,
	`Rand McNally`,
	`Rank and Filed`,
	`Ranku`,
	`Rapid Cycle Solutions`,
	`realtor.com`,
	`Recargo`,
	`ReciPal`,
	`Redfin`,
	`RedLaser`,
	`Reed Elsevier`,
	`REI Systems`,
	`Relationship Science`,
	`Remi`,
	`Retroficiency`,
	`Revaluate`,
	`Revelstone`,
	`Rezolve Group`,
	`Rivet Software`,
	`Roadify Transit`,
	`Robinson + Yu`,
	`Russell Investments`,
	`Sage Bionetworks`,
	`SAP`,
	`SAS`,
	`Scale Unlimited`,
	`Science Exchange`,
	`Seabourne`,
	`SeeClickFix`,
	`SigFig`,
	`Simple Energy`,
	`SimpleTuition`,
	`SlashDB`,
	`Smart Utility Systems`,
	`SmartAsset`,
	`SmartProcure`,
	`Smartronix`,
	`SnapSense`,
	`Social Explorer`,
	`Social Health Insights`,
	`SocialEffort`,
	`Socrata`,
	`Solar Census`,
	`SolarList`,
	`Sophic Systems Alliance`,
	`S&P Capital IQ`,
	`SpaceCurve`,
	`SpeSo Health`,
	`Spikes Cavell Analytic`,
	`Splunk`,
	`Spokeo`,
	`SpotCrime`,
	`SpotHero.com`,
	`Stamen Design`,
	`Standard and Poor's`,
	`State Farm Insurance`,
	`Sterling Infosystems`,
	`Stevens Worldwide Van Lines`,
	`STILLWATER SUPERCOMPUTING INC`,
	`StockSmart`,
	`Stormpulse`,
	`StreamLink Software`,
	`StreetCred Software,`,
	`StreetEasy`,
	`Suddath`,
	`Symcat`,
	`Synthicity`,
	`T. Rowe Price`,
	`Tableau Software`,
	`TagniFi`,
	`Telenav`,
	`Tendril`,
	`Teradata`,
	`The Advisory Board Company`,
	`The Bridgespan Group`,
	`The DocGraph Journal`,
	`The Govtech Fund`,
	`The Schork Report`,
	`The Vanguard Group`,
	`Think Computer`,
	`Thinknum`,
	`Thomson Reuters`,
	`TopCoder`,
	`TowerData`,
	`TransparaGov`,
	`TransUnion`,
	`TrialTrove`,
	`TrialX`,
	`Trintech`,
	`TrueCar`,
	`Trulia`,
	`TrustedID`,
	`TuvaLabs`,
	`Uber`,
	`Unigo`,
	`United Mayflower`,
	`Urban Airship`,
	`Urban Mapping,`,
	`US Green Data`,
	`U.S. News Schools`,
	`USAA Group`,
	`USSearch`,
	`Verdafero`,
	`Vimo`,
	`VisualDoD,`,
	`Vital Axiom | Niinja`,
	`VitalChek`,
	`Vitals`,
	`Vizzuality`,
	`Votizen`,
	`Walk Score`,
	`WaterSmart Software`,
	`WattzOn`,
	`Way Better Patents`,
	`Weather Channel`,
	`Weather Decision Technologies`,
	`Weather Underground`,
	`WebFilings`,
	`Webitects`,
	`WebMD`,
	`Weight Watchers`,
	`WeMakeItSafer`,
	`Wheaton World Wide Moving`,
	`Whitby Group`,
	`Wolfram Research`,
	`Wolters Kluwer`,
	`Workhands`,
	`Xatori`,
	`Xcential`,
	`xDayta`,
	`Xignite`,
	`Yahoo`,
	`Zebu Compliance Solutions`,
	`Yelp`,
	`YourMapper`,
	`Zillow`,
	`ZocDoc`,
	`Zonability`,
	`Zoner`,
	`Zurich Insurance`,
}