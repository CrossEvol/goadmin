SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (1, 'board', NULL);
INSERT INTO `category` VALUES (2, 'humour', NULL);
INSERT INTO `category` VALUES (3, 'sugar', NULL);
INSERT INTO `category` VALUES (4, 'bunch', NULL);
INSERT INTO `category` VALUES (5, 'covey', NULL);
INSERT INTO `category` VALUES (6, 'bag', 1);
INSERT INTO `category` VALUES (7, 'project', 1);
INSERT INTO `category` VALUES (8, 'cast', 1);
INSERT INTO `category` VALUES (9, 'life', 1);
INSERT INTO `category` VALUES (10, 'eye', 1);
INSERT INTO `category` VALUES (11, 'cast', 2);
INSERT INTO `category` VALUES (12, 'london', 2);
INSERT INTO `category` VALUES (13, 'city', 2);
INSERT INTO `category` VALUES (14, 'loneliness', 2);
INSERT INTO `category` VALUES (15, 'surgeon', 2);
INSERT INTO `category` VALUES (16, 'joy', 3);
INSERT INTO `category` VALUES (17, 'solitude', 3);
INSERT INTO `category` VALUES (18, 'turkey', 3);
INSERT INTO `category` VALUES (19, 'trend', 3);
INSERT INTO `category` VALUES (20, 'book', 3);
INSERT INTO `category` VALUES (21, 'caravan', 4);
INSERT INTO `category` VALUES (22, 'friendship', 4);
INSERT INTO `category` VALUES (23, 'set', 4);
INSERT INTO `category` VALUES (24, 'racism', 4);
INSERT INTO `category` VALUES (25, 'butter', 4);
INSERT INTO `category` VALUES (26, 'camp', 5);
INSERT INTO `category` VALUES (27, 'cackle', 5);
INSERT INTO `category` VALUES (28, 'bread', 5);
INSERT INTO `category` VALUES (29, 'fear', 5);
INSERT INTO `category` VALUES (30, 'dress', 5);

-- ----------------------------
-- Records of group
-- ----------------------------
INSERT INTO `group` VALUES (1, 'heels', 'Last those what were why.');
INSERT INTO `group` VALUES (2, 'scheme', 'Cry scold relaxation help our.');
INSERT INTO `group` VALUES (3, 'leap', 'Justice worrisome should am harvest.');
INSERT INTO `group` VALUES (4, 'murder', 'Later this thing then these.');
INSERT INTO `group` VALUES (5, 'part', 'Then consequently bravo wild Atlantean.');
INSERT INTO `group` VALUES (6, 'murder', 'There lovely courage hourly being.');
INSERT INTO `group` VALUES (7, 'shoulder', 'Cautious besides daily mine this.');
INSERT INTO `group` VALUES (8, 'sedge', 'That hard flour data aha.');
INSERT INTO `group` VALUES (9, 'quiver', 'This who fortnightly block there.');
INSERT INTO `group` VALUES (10, 'ream', 'Daily from accommodation naughty despite.');

-- ----------------------------
-- Records of todo
-- ----------------------------
INSERT INTO `todo` VALUES (1, 'Othello', 3, 3.35489, 'PENDING', '2024-04-01 14:07:28', '2024-04-01 14:07:28', '2024-04-01 19:06:29', 'HIGH', 'Group exaltation this his safety wisp including yours other sedge.', 4);
INSERT INTO `todo` VALUES (2, 'The Idiot', 6, 8.40533, 'PENDING', '2024-04-01 14:07:28', '2024-04-01 14:07:28', '2024-04-01 15:06:29', 'MEDIUM', 'Little despite hall week also idea his where i.e. herself.', 10);
INSERT INTO `todo` VALUES (3, 'The Old Man and the Sea', 8, 1.21135, 'PENDING', '2024-04-01 14:07:28', '2024-04-01 14:07:28', '2024-04-01 20:06:29', 'MEDIUM', 'It so mine lastly me certain that everything band how.', 12);
INSERT INTO `todo` VALUES (4, 'Things Fall Apart', 9, 9.76061, 'PENDING', '2024-04-01 14:07:28', '2024-04-01 14:07:28', '2024-04-01 23:06:29', 'MEDIUM', 'Twist of pod your include is downstairs being selfish in.', 26);
INSERT INTO `todo` VALUES (5, 'The Golden Notebook', 6, 5.43767, 'PENDING', '2024-04-01 14:07:28', '2024-04-01 14:07:28', '2024-04-01 21:06:29', 'MEDIUM', 'Who it though group yesterday anyone group several that want.', 8);
INSERT INTO `todo` VALUES (6, 'Wuthering Heights', 2, 8.41354, 'PENDING', '2024-04-01 14:07:28', '2024-04-01 14:07:28', '2024-04-01 18:06:29', 'MEDIUM', 'Everything her are any aircraft you jump theirs flick formerly.', 7);
INSERT INTO `todo` VALUES (7, 'Leaves of Grass', 7, 3.41801, 'PENDING', '2024-04-01 14:07:28', '2024-04-01 14:07:28', '2024-04-02 02:06:29', 'MEDIUM', 'Outside patiently also one our that my then road eek.', 9);
INSERT INTO `todo` VALUES (8, 'Othello', 1, 5.10213, 'PENDING', '2024-04-01 14:07:28', '2024-04-01 14:07:28', '2024-04-01 20:06:29', 'MEDIUM', 'Timing have these rarely upon wait what cackle he chair.', 14);
INSERT INTO `todo` VALUES (9, 'Madame Bovary', 1, 5.16583, 'PENDING', '2024-04-01 14:07:28', '2024-04-01 14:07:28', '2024-04-01 23:06:29', 'MEDIUM', 'Crowd today just words wow doctor at I under swing.', 3);
INSERT INTO `todo` VALUES (10, 'Ulysses', 2, 7.74204, 'PENDING', '2024-04-01 14:07:28', '2024-04-01 14:07:28', '2024-04-01 23:06:29', 'MEDIUM', 'These panther it can these next ever later which Orwellian.', 6);

-- ----------------------------
-- Records of todo_detail
-- ----------------------------
INSERT INTO `todo_detail` VALUES (1, 'Ride your your just let eek slavery frail yours his.', 'http://www.investore-markets.io/b2c/one-to-one/synergistic', 1);
INSERT INTO `todo_detail` VALUES (2, 'Pad fuel what single brilliance her furthermore whose east grieving.', 'http://www.corporatesystems.org/integrated/wireless/dynamic/initiatives', 2);
INSERT INTO `todo_detail` VALUES (3, 'The than one hand throw always sail none late youth.', 'http://www.forwardsystems.name/solutions/viral', 3);
INSERT INTO `todo_detail` VALUES (4, 'Then table so neither weekly cloud someone band woman straw.', 'https://www.productweb-enabled.info/relationships', 4);
INSERT INTO `todo_detail` VALUES (5, 'Information American ouch whom walk utterly she a tonight bunch.', 'https://www.legacyvisionary.biz/world-class', 5);
INSERT INTO `todo_detail` VALUES (6, 'Out her these whoa covey it wisdom that luck how.', 'http://www.centralopen-source.io/compelling/architect/roi', 6);
INSERT INTO `todo_detail` VALUES (7, 'Early though ream sit someone cook bundle finish everyone whoever.', 'http://www.dynamicvirtual.io/utilize/ubiquitous/compelling/enable', 7);
INSERT INTO `todo_detail` VALUES (8, 'Someone my many that wade mine be to admit on.', 'https://www.seniorrobust.com/e-commerce/transition/infrastructures/mindshare', 8);
INSERT INTO `todo_detail` VALUES (9, 'Contrast had yearly he that regiment towards English anyway straightaway.', 'https://www.globalrevolutionize.name/disintermediate/repurpose/convergence/proactive', 9);
INSERT INTO `todo_detail` VALUES (10, 'Anyone his aha let annually nobody freedom of few motivation.', 'https://www.seniorcustomized.name/holistic', 10);

-- ----------------------------
-- Records of todo_tag
-- ----------------------------
INSERT INTO `todo_tag` VALUES (1, 'AnxiousSardine', 1, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (2, 'SleepyTie', 1, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (3, 'DisturbedStreet', 1, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (4, 'DamsonLong7', 1, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (5, 'FilthyArmadillo', 1, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (6, 'PhysalisSpotted028', 2, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (7, 'KumquatElated269', 2, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (8, 'BookstoreFlyer', 2, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (9, 'LemonyJaw', 2, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (10, 'LemonyComputer', 2, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (11, 'TameCardigan6', 3, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (12, 'ToyShakeer', 3, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (13, 'DonkeyRideer', 3, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (14, 'DamsonAdorable8', 3, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (15, 'KumquatHorrible52', 3, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (16, 'SpottedWeasel0', 4, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (17, 'WatermelonEncouraging', 4, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (18, 'WearyGerbil', 4, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (19, 'LipsSnoreer9', 4, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (20, 'PitayaBlushing592', 4, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (21, 'ChickenStander', 5, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (22, 'StrawberryElated', 5, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (23, 'PitayaGleaming', 5, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (24, 'FoolishBank', 5, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (25, 'ThankfulSpoon', 5, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (26, 'ShyPlatypus', 6, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (27, 'AnnoyingTrenchCoat', 6, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (28, 'WashingMachineFlyer', 6, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (29, 'GracefulRainbow0', 6, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (30, 'DisturbedLizard986', 6, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (31, 'NervousPig', 7, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (32, 'GiftSkier', 7, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (33, 'CharmingBee', 7, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (34, 'NurseDanceer688', 7, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (35, 'PerfectVoice', 7, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (36, 'CautiousPlatypus', 8, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (37, 'ChokerClaper', 8, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (38, 'BackShakeer', 8, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (39, 'MagnificentRhinoceros177', 8, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (40, 'PineappleWandering', 8, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (41, 'JoyousCrab473', 9, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (42, 'YellowFly0', 9, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (43, 'ExpensiveEel5', 9, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (44, 'EncouragingMango', 9, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (45, 'GlamorousMonkey', 9, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (46, 'StrangeDeer178', 10, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (47, 'FeijoaClumsy', 10, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (48, 'PastaCooker', 10, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (49, 'CleverApple341', 10, '2024-04-01 14:07:28');
INSERT INTO `todo_tag` VALUES (50, 'LonelyPhone20', 10, '2024-04-01 14:07:28');

-- ----------------------------
-- Records of todosongroups
-- ----------------------------
INSERT INTO `todosongroups` VALUES (1, 1, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (1, 2, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (1, 3, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (1, 4, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (1, 5, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (1, 6, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (1, 7, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (1, 8, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (1, 9, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (1, 10, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (2, 1, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (2, 2, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (2, 3, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (2, 4, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (2, 5, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (2, 6, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (2, 7, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (2, 8, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (2, 9, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (3, 1, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (3, 2, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (3, 3, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (3, 4, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (3, 5, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (3, 6, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (3, 7, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (3, 8, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (4, 1, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (4, 2, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (4, 3, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (4, 4, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (4, 5, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (4, 6, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (4, 7, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (5, 1, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (5, 2, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (5, 3, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (5, 4, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (5, 5, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (5, 6, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (6, 1, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (6, 2, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (6, 3, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (6, 4, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (6, 5, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (7, 1, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (7, 2, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (7, 3, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (7, 4, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (8, 1, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (8, 2, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (8, 3, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (9, 1, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (9, 2, '2024-04-01 14:07:28');
INSERT INTO `todosongroups` VALUES (10, 1, '2024-04-01 14:07:28');

SET FOREIGN_KEY_CHECKS = 1;
