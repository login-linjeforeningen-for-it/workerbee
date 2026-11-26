------------------
-- Dummy Data
------------------

INSERT INTO "quotes" ("author", "quoted", "content") VALUES
('Hermanius Elganius', 'Gjermund', 'Han er en innvikler.'),
('Ole', 'Hoff', '9090 incase 5432 is in use.'),
('Eirik', 'Riwa','Jeg blir fysisk sint av å se på kode.'),
('Ole', 'Hoff', 'Eg tenker en pepperbiff, med flødegratinerde poteter og stekegrad på rååååååååååå');

INSERT INTO "alerts" (service, page, title_en, title_no, description_en, description_no) VALUES
('beehive', '/events', 'Important Beehive Update', 'Viktig Beehive Oppdatering', 'This is an important message for Beehive users. Please read carefully.', 'Dette er en viktig melding for Beehive brukere. Vennligst les nøye.'),
('beehive', '/jobs', 'New Job Opportunities', 'Nye Jobbmuligheter', 'Check out the latest job openings available for members.', 'Sjekk ut de siste ledige stillingene tilgjengelig for medlemmer.'),
('tekkom', '/events', 'TekKom Event Announcement', 'TekKom Arrangement Annonsering', 'Join us for upcoming TekKom events and workshops.', 'Bli med oss for kommende TekKom arrangementer og workshops.'),
('tekkom', '/jobs', 'TekKom Job Listings', 'TekKom Jobbannonser', 'Explore new job listings through TekKom.', 'Utforsk nye jobbannonser via TekKom.');

INSERT INTO "albums" ("name_no", "name_en", "description_no", "description_en", "year", "event_id", "created_at", "updated_at") VALUES
('Album', 'Album', 'Photos from events', 'Bilder fra arrangementer', 2023, NULL, now(), now()),
('Album', 'Album', 'Photos from events', 'Bilder fra arrangementer', 2024, NULL, now(), now());

INSERT INTO "audiences" ("name_en", "name_no")
VALUES
('Login', 'Login'),
('Active members', 'Aktive medlemmer'),
('Students', 'Studenter'),
('Open', 'Åpen'),
('Bachelor', 'Bachelor'),
('Master', 'Master'),
('PhD', 'PhD'),
('First semester', 'Første semester'),
('Second semester', 'Andre semester'),
('Third semester', 'Tredje semester'),
('Fourth semester', 'Fjerde semester'),
('Fifth semester', 'Femte semester'),
('Sixth semester', 'Sjette semester'),
('Seventh semester', 'Sjuende semester');

INSERT INTO "skills" ("name") VALUES
('Programmering'),
('Datavisualisering'),
('Prosjektledelse'),
('Markedsføring'),
('Kundebehandling');

INSERT INTO "cities" ("name") VALUES
('Arendal'),
('Bodø'),
('Drammen'),
('Fredrikstad'),
('Gjøvik'),
('Halden'),
('Hamar'),
('Harstad'),
('Haugesund'),
('Kristiansand'),
('Kristiansund'),
('Larvik'),
('Lillehammer'),
('Mo i Rana'),
('Molde'),
('Moss'),
('Porsgrunn'),
('Sandefjord'),
('Sandnes'),
('Sarpsborg'),
('Skien'),
('Steinkjer'),
('Tønsberg'),
('Ålesund'),
('Alta'),
('Oslo'),
('Trondheim'),
('Bergen'),
('Stavanger'),
('Tromsø');

INSERT INTO "honey" ("service", "language", "page", "text") VALUES
('beehive', 'no', '/events', '{"title": "Arrangementer", "description": "Alle kommende arrangementer for Login."}'),
('beehive', 'en', '/events', '{"title": "Events", "description": "All upcoming events for Login."}'),
('beehive', 'no', '/jobs', '{"title": "Jobber", "description": "Ledige stillinger og verv for medlemmer."}'),
('beehive', 'en', '/jobs', '{"title": "Jobs", "description": "Open positions and roles for members."}'),
('tekkom', 'no', '/events', '{"title": "TekKom Arrangementer", "description": "TekKom sine arrangementer og workshops."}'),
('tekkom', 'en', '/events', '{"title": "TekKom Events", "description": "TekKom events and workshops."}'),
('tekkom', 'no', '/jobs', '{"title": "TekKom Jobber", "description": "Jobbmuligheter via TekKom."}'),
('tekkom', 'en', '/jobs', '{"title": "TekKom Jobs", "description": "Job opportunities via TekKom."}');


INSERT INTO "locations" (
  "name_no", 
  "name_en", 
  "type", 
  "mazemap_campus_id", 
  "mazemap_poi_id", 
  "address_street", 
  "address_postcode", 
  "city_id", 
  "coordinate_lat", 
  "coordinate_lon", 
  "url", 
  "updated_at", 
  "created_at"
)
VALUES
('Universitetet i Oslo', 'University of Oslo', 'address', NULL, NULL, 
  'Problemveien 7', 313, 26, 59.9390, 10.7205, NULL, now(), now()),
('Nidarosdomen', 'Nidaros Cathedral', 'coords', NULL, NULL, 
  NULL, NULL, 27, 63.4277, 10.3969, NULL, now(), now()),
('Bergenhus Festning', 'Bergenhus Fortress', 'address', NULL, NULL, 
  'Bergenhus', 5003, 28, 60.3993, 5.3221, NULL, now(), now()),
('Stavanger Forum', 'Stavanger Forum', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 1', 'Bestegata 1', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 2', 'Bestegata 2', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 3', 'Bestegata 3', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 4', 'Bestegata 4', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 5', 'Bestegata 5', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 6', 'Bestegata 6', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 7', 'Bestegata 7', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 8', 'Bestegata 8', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 9', 'Bestegata 9', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 10', 'Bestegata 10', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 11', 'Bestegata 11', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 12', 'Bestegata 12', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 13', 'Bestegata 13', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 14', 'Bestegata 14', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 15', 'Bestegata 15', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 16', 'Bestegata 16', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 17', 'Bestegata 17', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 18', 'Bestegata 18', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 19', 'Bestegata 19', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 20', 'Bestegata 20', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Tromsø Bibliotek', 'Tromsø Library', 'address', NULL, NULL, 
  'Grønnegata 94', 9008, 30, 69.6489, 18.9551, NULL, now(), now());
INSERT INTO "rules" (
  "name_no", "name_en", "description_no", "description_en", "updated_at", "created_at"
)
VALUES
('Ingen eksterne matvarer', 'No External Food Allowed', 
  'Det er ikke tillatt å ta med matvarer utenfra til arrangementet.', 
  'Bringing external food to the event is not allowed.', 
  now(), now()),
('Påmelding er obligatorisk', 'Registration Required', 
  'Alle deltakere må registrere seg på forhånd for å delta.', 
  'All part_timeicipants must register in advance to attend.', 
  now(), now()),
('Stille mobiltelefoner', 'Silence Mobile Phones', 
  'Alle mobiltelefoner må settes på lydløs under arrangementet.', 
  'All mobile phones must be set to silent mode during the event.', 
  now(), now()),
('Ikke røyk', 'No Smoking', 
  'Røyking er ikke tillatt i eller rundt arrangementsområdet.', 
  'Smoking is not permitted in or around the event area.', 
  now(), now()),
('Gyldig billett kreves', 'Valid Ticket Required', 
  'Alle deltakere må ha en gyldig billett for å delta på arrangementet.', 
  'All attendees must present a valid ticket to part_timeicipate in the event.', 
  now(), now());

INSERT INTO "organizations" (
  "name_no", 
  "name_en", 
  "description_no", 
  "description_en", 
  "link_homepage", 
  "link_linkedin", 
  "link_facebook", 
  "link_instagram", 
  "logo", 
  "updated_at", 
  "created_at"
)
VALUES
  ('Universitetet i Oslo', 'University of Oslo', 
   'Universitetet i Oslo er Norges største universitet, med et bredt fagtilbud.', 
   'The University of Oslo is Norways largest university, offering a wide range of programs.', 
   'https://www.uio.no', 'https://www.linkedin.com/school/university-of-oslo', 
   'https://www.facebook.com/uni.oslo', 'https://www.instagram.com/uniofoslo', 
   'logo.png', now(), now()),
  ('Norges teknisk-naturvitenskapelige universitet', 'Norwegian University of Science and Technology', 
   'NTNU er et teknisk universitet i Trondheim, kjent for sin forskning på teknologi og naturvitenskap.', 
   'NTNU is a technical university in Trondheim, known for its research in technology and natural sciences.', 
   'https://www.ntnu.no', 'https://www.linkedin.com/school/ntnu', 
   'https://www.facebook.com/NTNU.no', 'https://www.instagram.com/ntnu_official', 
   'logo.png', now(), now()),
  ('DNB ASA', 'DNB ASA', 
   'DNB er Norges største finanskonsern med et bredt tilbud av finansielle tjenester.', 
   'DNB is Norways largest financial group, offering a wide range of financial services.', 
   'https://www.dnb.no', 'https://www.linkedin.com/company/dnb', 
   'https://www.facebook.com/dnb.no', 'https://www.instagram.com/dnb.no', 
   'logo.png', now(), now()),
  ('Telenor ASA', 'Telenor ASA', 
   'Telenor er et ledende teleselskap som tilbyr mobil- og bredbåndstjenester.', 
   'Telenor is a leading telecommunications company offering mobile and broadband services.', 
   'https://www.telenor.no', 'https://www.linkedin.com/company/telenor', 
   'https://www.facebook.com/telenor', 'https://www.instagram.com/telenor', 
   'logo.png', now(), now()),
  ('SINTEF', 'SINTEF', 
   'SINTEF er en av Europas største uavhengige forskningsorganisasjoner, kjent for sitt arbeid innen teknologi og innovasjon.', 
   'SINTEF is one of Europes largest independent research organizations, known for its work in technology and innovation.', 
   'https://www.sintef.no', 'https://www.linkedin.com/company/sintef', 
   'https://www.facebook.com/SINTEF', 'https://www.instagram.com/sintef', 
   'logo.png', now(), now());

INSERT INTO "categories" ("name_en", "name_no", "color")
VALUES
('Login', 'Login', '#fd8738'),
('TekKom', 'TekKom', '#a206c9'),
('CTFKom', 'CTFKom', '#2da62b'),
('EvntKom', 'EvntKom', '#d62f43'),
('PR', 'PR', '#ffff00'),
('BedKom', 'BedKom', '#1f56c5'),
('SatKom', 'SatKom', '#64ddd7'),
('BroomBroom', 'BroomBroom', '#cd53d8'),
('Pearlgroup', 'Perlegruppa', '#cd53d8'),
('Bookclub', 'Bokklubb', '#cd53d8'),
('Houseband', 'Husbandet', '#cd53d8'),
('Social', 'Sosialt', '#d62f43'),
('Cyberdays', 'Cyberdagene', 'linear-gradient(120deg, hsla(217, 100%, 50%, 1) 10%, hsla(186, 100%, 69%, 1) 100%)'),
('Buddyweek', 'Fadderuka', '#fa75a6'),
('BarKom', 'BarKom', '#f1c40f'),
('Other', 'Andre', '#545b5f');

INSERT INTO "jobs" (
  "visible", 
  "highlight", 
  "title_no", 
  "title_en", 
  "position_title_no", 
  "position_title_en", 
  "description_short_no", 
  "description_short_en", 
  "description_long_no", 
  "description_long_en", 
  "job_type_id", 
  "time_publish", 
  "time_expire", 
  "banner_image", 
  "organization_id", 
  "application_url", 
  "updated_at", 
  "created_at"
)
VALUES
(true, false, 'Softwareutvikler', 'Software Developer', 'Junior utvikler', 
  'Junior Developer', 'En spennende mulighet for nyutdannede til å utvikle programvare.', 
  'An exciting opportunity for recent graduates to develop software.', 
  'Som Junior Software Developer vil du være med på utvikling av applikasjoner og programvare.', 
  'As a Junior Software Developer, you will be involved in the development of applications and software.', 
  1, now() + interval '1 days', now() + interval '1 years','https://www.example.com/banner.jpg', 
  1, 'https://www.uio.no/job-apply', now(), now() - interval '4 days'),
(true, true, 'Markedsføringskoordinator', 'Marketing Coordinator', 'Markedsføringsspesialist', 
  'Marketing Specialist', 'Bli en del av vårt markedsføringsteam og jobb med spennende prosjekter.', 
  'Join our marketing team and work on exciting projects.', 
  'Som markedsføringskoordinator vil du ha ansvar for markedsføring og kommunikasjon på tvers av kanaler.', 
  'As a Marketing Coordinator, you will be responsible for marketing and communication across channels.', 
  2, now() + interval '1 days', now() + interval '1 years', 'https://www.example.com/banner2.jpg', 
  3, 'https://www.dnb.no/job-apply', now(), now()),
(false, false, 'Prosjektleder', 'Project Manager', 'Senior prosjektleder', 
  'Senior Project Manager', 'Vi søker en erfaren prosjektleder til å lede store prosjekter.', 
  'We are looking for an experienced project manager to lead large projects.', 
  'Som prosjektleder vil du ha ansvar for å lede prosjekter fra start til slutt, inkludert budsjett og tidsplanlegging.', 
  'As a Project Manager, you will be responsible for leading projects from start to finish, including budgeting and scheduling.', 
  3, now() + interval '1 days', now() + interval '1 years', 'https://www.example.com/banner3.jpg', 
  4, 'https://www.telenor.no/job-apply', now(), now()),
(true, false, 'Kundestøtteagent', 'Customer Support Agent', 'Kundestøtteansvarlig', 
  'Customer Support Manager', 'Bli en del av vårt kundeserviceteam og hjelp kunder med deres henvendelser.', 
  'Join our customer service team and assist customers with their inquiries.', 
  'Som kundestøtteansvarlig vil du hjelpe kunder via telefon, e-post og chat, samt sikre god kundetilfredshet.', 
  'As a Customer Support Manager, you will assist customers via phone, email, and chat, ensuring high customer satisfaction.', 
  2, now() + interval '1 days', now() + interval '1 years', 'https://www.example.com/banner4.jpg', 
  5, 'https://www.sintef.no/job-apply', now(), now()),
(true, true, 'Dataanalytiker', 'Data Analyst', 'Dataanalytiker', 
  'Data Analyst', 'Er du en dataanalytiker som elsker å finne innsikt fra store datamengder?', 
  'Are you a data analyst who loves to derive insights from large datasets?', 
  'Som dataanalytiker vil du analysere data for å identifisere trender og lage rapporter som støtter beslutningstaking.', 
  'As a Data Analyst, you will analyze data to identify trends and create reports that support decision-making.', 
  1, now() + interval '1 days', now() + interval '1 years', 'https://www.example.com/banner5.jpg', 
  2, 'https://www.ntnu.no/job-apply', now(), now());

INSERT INTO "ad_skill_relation" ("job_id", "skill_id")
VALUES
(1, 1), -- Software Developer (UiO) - Programmering
(1, 2), -- Software Developer (UiO) - Datavisualisering
(2, 3), -- Marketing Coordinator (DNB) - Prosjektledelse
(3, 4), -- Project Manager (Telenor) - Markedsføring
(3, 5); -- Project Manager (Telenor) - Kundebehandling

INSERT INTO "ad_city_relation" ("job_id", "city_id")
VALUES
(1, 26),  -- Software Developer (UiO) - Oslo
(2, 29),  -- Marketing Coordinator (DNB) - Stavanger
(3, 30),  -- Project Manager (Telenor) - Tromsø
(4, 27),  -- Customer Support Agent (SINTEF) - Trondheim
(5, 27);  -- Data Analyst (NTNU) - Trondheim

INSERT INTO "events" (
  "visible", "name_no", "name_en", "description_no", "description_en", 
  "informational_no", "informational_en", "time_type", "time_start", "time_end", 
  "time_publish", "time_signup_release", "time_signup_deadline", "canceled", 
  "digital", "highlight", "image_small", "image_banner", "link_facebook", 
  "link_discord", "link_signup", "link_stream", "capacity", "is_full", 
  "organization_id", "location_id", "parent_id", "rule_id", "audience_id", "category_id", "created_at", "updated_at"
)
VALUES
-- Event 1: Published, starts in 3 days (WILL show in public API)
(true, 'Hackathon Oslo', 'Hackathon Oslo', 
 'Bli med på en spennende hackathon i Oslo!', 'Join an exciting hackathon in Oslo!', 
 'Mer informasjon kommer snart.', 'More information coming soon.',
 'whole_day', NOW() + interval '3 days', NOW() + interval '3 days' + interval '9 hours', 
 NOW() - interval '15 days', NOW() - interval '10 days', NOW() + interval '2 days', 
 false, true, false, NULL, 'https://www.example.com/banner_hackathon.jpg', 
 NULL, NULL, NULL, NULL, 100, false, 1, 1, NULL, 1, 1, 1, NOW() - interval '3 days', NOW()),

-- Event 2: Published, starts in 2 weeks (WILL show in public API)
(true, 'Tech Conference Bergen', 'Tech Conference Bergen', 
 'Lær om de nyeste teknologiene på Tech Conference i Bergen.', 
 'Learn about the latest technologies at Tech Conference in Bergen.', 
 'Påmelding nødvendig.', 'Registration required.',
 'whole_day', NOW() + interval '14 days', NOW() + interval '14 days' + interval '8 hours', 
 NOW() - interval '5 days', NOW() - interval '3 days', NOW() + interval '12 days', 
 false, true, true, NULL, 'https://www.example.com/banner_tech_conference.jpg', 
 NULL, NULL, NULL, NULL, 200, false, 2, 2, NULL, 1, 2, 1, NOW(), NOW()),

-- Event 3: NOT YET published (will NOT show in public API - publishes in 1 week)
(true, 'AI Workshop Trondheim', 'AI Workshop Trondheim', 
 'Utforsk kunstig intelligens i Trondheim!', 'Explore artificial intelligence in Trondheim!', 
 'Gratis workshop for alle interesserte.', 'Free workshop for all interested.',
 'whole_day', NOW() + interval '21 days', NOW() + interval '21 days' + interval '6 hours', 
 NOW() + interval '7 days', NOW() + interval '7 days', NOW() + interval '20 days', 
 false, true, false, NULL, 'https://www.example.com/banner_ai_workshop.jpg', 
 NULL, NULL, NULL, NULL, 50, false, 3, 3, NULL, 1, 1, 1, NOW(), NOW()),

-- Event 4: Published, HAPPENING NOW (started 1 hour ago, ends in 3 hours - WILL show in public API)
(true, 'Cybersecurity Summit Stavanger', 'Cybersecurity Summit Stavanger', 
 'Lær om cybersikkerhet i Stavanger!', 'Learn about cybersecurity in Stavanger!', 
 'Fokus på praktiske løsninger.', 'Focus on practical solutions.',
 'default', NOW() - interval '1 hour', NOW() + interval '3 hours', 
 NOW() - interval '20 days', NOW() - interval '15 days', NOW() - interval '2 days', 
 false, true, true, NULL, 'https://www.example.com/banner_cybersecurity_summit.jpg', 
 NULL, NULL, NULL, NULL, 150, false, 4, 3, NULL, 3, 1, 4, NOW(), NOW()),

-- Event 5: Published, starts in 1 month (WILL show in public API)
(true, 'Cloud Computing Meetup Tromsø', 'Cloud Computing Meetup Tromsø', 
 'Møt eksperter innen skyteknologi i Tromsø.', 'Meet cloud technology experts in Tromsø.', 
 'Networking muligheter.', 'Networking opportunities.',
 'whole_day', NOW() + interval '30 days', NOW() + interval '30 days' + interval '4 hours', 
 NOW() - interval '2 days', NOW() + interval '5 days', NOW() + interval '28 days', 
 false, true, false, NULL, 'https://www.example.com/banner_cloud_meetup.jpg', 
 NULL, NULL, NULL, NULL, 80, false, 5, 1, NULL, 2, 2, 5, NOW(), NOW());