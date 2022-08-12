package database

var (
	danmaku_report = `
					CREATE TABLE IF NOT EXISTS danmaku_report (
						cid int(8) PRIMARY KEY NOT NULL,	-- 弹幕ID
						id varchar(128) NOT NULL,			-- 弹幕池id
						text varchar(128) NOT NULL, 		-- 举报内容
						type varchar(128) NOT NULL, 		-- 举报类型
						time varchar(128) NOT NULL, 		-- 举报时间
						ip varchar(12) NOT NULL 			-- 发送弹幕的IP地址
	  				);
					`
	danmaku_ip = `
				CREATE TABLE IF NOT EXISTS danmaku_ip (
				ip varchar(12) PRIMARY KEY NOT NULL, 		-- 发送弹幕的IP地址,
				c int(1) NOT NULL DEFAULT '1', 				-- 规定时间内的发送次数
				time int(10) NOT NULL
	  			);
	`

	danmaku_list = `
					CREATE TABLE IF NOT EXISTS danmaku_list (
					id varchar(32) NOT NULL, 				-- 弹幕池id
					cid INTEGER PRIMARY KEY AUTOINCREMENT, 	-- 弹幕id
					type varchar(128) NOT NULL, 			-- 弹幕类型
					text varchar(128) NOT NULL,				-- 弹幕内容
					color varchar(128) NOT NULL,			-- 弹幕颜色
					size varchar(128) NOT NULL,				-- 弹幕大小
					videotime float(24,3) NOT NULL,			-- 时间点
					ip varchar(128) NOT NULL,				-- 用户ip
					time int(10) NOT NULL					-- 发送时间
	  )
	`
)
