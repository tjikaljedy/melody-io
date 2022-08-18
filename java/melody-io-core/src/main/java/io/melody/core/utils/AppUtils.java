package io.melody.core.utils;

import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.TimeZone;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class AppUtils {
	private static final Logger log = LoggerFactory.getLogger(AppUtils.class);
	public static final String DATE_PATTERN = "yyyy-MM-dd";
	public static final String GA_DATE_PATTERN = "yyyyMMdd";
	public static final String MONGO_DATE_PATTERN = "EEE MMM dd HH:mm:ss Z yyyy";
	public static final String HUMAN_DATE_PATTERN = "dd MMM yyy";
	public static final String DEFAULT_DATE_PATTERN = "yyyy-MM-dd HH:mm:ss";
	public static final String ZONE_UTC = "UTC";
	public static final String ZONE_GMT_PLUS7 = "GMT+7:00";
	public static final String ZONE_GMT_PLUS8 = "GMT+8:00";
	public static final String RFC3339 = "yyyy-MM-dd'T'HH:mm:ssZ";
	
	public static String dateToDateStringFormat(Date date, String dateFormat, String timeZone) {
		SimpleDateFormat df = new SimpleDateFormat(dateFormat);
		df.setTimeZone(TimeZone.getTimeZone(timeZone));
		return df.format(date);
	}
}
